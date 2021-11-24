package showtable

import (
	paddle "github.com/PaddleHQ/go-aws-ssm"
	albv2 "github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	ecs "github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	ecs_patterns "github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateStackStackProps struct {
	cdk.StackProps
}


type  RuntimePlatform struct {
	CPUArchitecture       string `json:"CpuArchitecture"`
	OperatingSystemFamily string `json:"OperatingSystemFamily"`
}


var pmstore *paddle.ParameterStore

func init() {
	var err error
	pmstore, err = paddle.NewParameterStore()
	if err != nil {
		panic("Cant connect to Parameter Store")
	}
	
	
}

func FargateStack(scope constructs.Construct, id string, props *FargateStackStackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)
	
	vpcid, err := pmstore.GetParameter("/network/basevpc", false)
	if err != nil {
		panic("Cant connect to Parameter Store")
	}
	vpc := ec2.Vpc_FromLookup(stack, aws.String("vpc"), &ec2.VpcLookupOptions{
		IsDefault:          aws.Bool(false),
		VpcId:              vpcid.Value,
	})
	cluster := ecs.NewCluster(stack, aws.String("ALBFargoECSCluster"), &ecs.ClusterProps{
		Vpc: vpc,
	})
	
	var actions = []*string{ 
		aws.String("dynamodb:Scan"),
		aws.String("ssm:GetParameter"),
	}

	var resources = []*string{
		 aws.String("*"),
	}

	role := iam.NewRole(stack, aws.String("fargateExecutionRole"), &iam.RoleProps{
		AssumedBy:           iam.NewServicePrincipal(aws.String("ecs-tasks.amazonaws.com"),
		 &iam.ServicePrincipalOpts{ Region: aws.String("eu-central-1")}),
		Description:         aws.String("Role for showtable"),
		InlinePolicies:      &map[string]iam.PolicyDocument{
			"showtable" : iam.NewPolicyDocument( &iam.PolicyDocumentProps{
				AssignSids: aws.Bool(true),
				Statements: &[]iam.PolicyStatement{
					iam.NewPolicyStatement( &iam.PolicyStatementProps{
						Actions:       &actions,
						Effect:        iam.Effect_ALLOW,
						Resources:     &resources,
					}),
				},
			},
		),
		},
		RoleName:            aws.String("showtable-gin"),
	})
	service := ecs_patterns.NewApplicationLoadBalancedFargateService(stack, aws.String("ALBFargoService"), &ecs_patterns.ApplicationLoadBalancedFargateServiceProps{
		Cluster: cluster,
		DesiredCount: aws.Float64(1),
		TaskImageOptions: &ecs_patterns.ApplicationLoadBalancedTaskImageOptions{
			Image: ecs.ContainerImage_FromAsset(aws.String("dist"),
				&ecs.AssetImageProps{
			
			}),
			TaskRole: role,
			ContainerPort: aws.Float64(8080),
		},
		PublicLoadBalancer: aws.Bool(true),
		Cpu: aws.Float64(256),
		MemoryLimitMiB: aws.Float64(512),
	})

	armPlatform := &RuntimePlatform{
		CPUArchitecture:       "ARM64",
		OperatingSystemFamily: "LINUX",
	}
	taskDefinition := service.TaskDefinition()
	// taskDefinition.Node()
	// Default child is IConstruct
	// ecs.CfnTaskDefinition(service.TaskDefinition().Node().DefaultChild())

	networkMode := string(service.TaskDefinition().NetworkMode())
	ecs.NewCfnTaskDefinition(stack, aws.String("ArmBasedTaskdefinition"), &ecs.CfnTaskDefinitionProps{
		ContainerDefinitions:    taskDefinition.Containers(),
		Cpu:                     aws.String("256"),
		Family:                  service.TaskDefinition().Family(),
		Memory:                  aws.String("512"),
		NetworkMode:             &networkMode,
		RequiresCompatibilities: &[]*string{ aws.String("FARGATE")},
		RuntimePlatform:         armPlatform,
		TaskRoleArn:             role.RoleArn(),
	})

	// 401 Unauthorized
	stethoscope := albv2.HealthCheck{
		Enabled:                 aws.Bool(true),
		HealthyHttpCodes:        aws.String("200,401"),
	}

	service.TargetGroup().ConfigureHealthCheck(&stethoscope)
	
	
	
	cdk.NewCfnOutput(stack, jsii.String("LoadBalancerDNS"), &cdk.CfnOutputProps{Value: service.LoadBalancer().LoadBalancerDnsName()})

	return stack
}

