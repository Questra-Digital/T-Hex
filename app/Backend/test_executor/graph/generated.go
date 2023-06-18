

package graph

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"test_executor/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)


func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Mutation struct {
		CloneGitRepository     func(childComplexity int, input model.CloneGitRepInput) int
		CreateDockerFile       func(childComplexity int, input model.CreateDockerFileInput) int
		DestroyAll             func(childComplexity int) int
		InitiateTest           func(childComplexity int) int
		ProjectBuildImage      func(childComplexity int, input model.StartTestInput) int
		ProjectContainerRemove func(childComplexity int, input model.ProjectContainerEndInput) int
		ProjectContainerStart  func(childComplexity int, input model.StartTestInput) int
		ProjectContainerStop   func(childComplexity int, input model.ProjectContainerEndInput) int
		SaveTestResult         func(childComplexity int, input model.SaveTestResultInput) int
		SeleniumPull           func(childComplexity int) int
		SeleniumRemove         func(childComplexity int, input model.SeleniumInput) int
		SeleniumStart          func(childComplexity int, input model.StartTestInput) int
		SeleniumStop           func(childComplexity int, input model.SeleniumInput) int
		StartTest              func(childComplexity int, input model.StartTestInput) int
		UpdateTest             func(childComplexity int, input model.UpdateTestInput) int
	}

	Projectinfos struct {
		GitEmail            func(childComplexity int) int
		GitLanguage         func(childComplexity int) int
		GitNoOfTestCases    func(childComplexity int) int
		GitProjectName      func(childComplexity int) int
		GitTestCaseFileName func(childComplexity int) int
		ID                  func(childComplexity int) int
	}

	Query struct {
		ProjectInfo  func(childComplexity int, id int) int
		ProjectInfos func(childComplexity int) int
		Setting      func(childComplexity int, id int) int
		Settings     func(childComplexity int) int
		Test         func(childComplexity int, id int) int
		Tests        func(childComplexity int) int
	}

	Settings struct {
		Browser               func(childComplexity int) int
		EnableLogs            func(childComplexity int) int
		ID                    func(childComplexity int) int
		NumberOfParallelTests func(childComplexity int) int
		Parallelism           func(childComplexity int) int
		StepByStepDebugging   func(childComplexity int) int
		Version               func(childComplexity int) int
	}

	Test struct {
		Duration func(childComplexity int) int
		ID       func(childComplexity int) int
		Project  func(childComplexity int) int
		Status   func(childComplexity int) int
		Testfile func(childComplexity int) int
		Type     func(childComplexity int) int
		Urlid    func(childComplexity int) int
		Username func(childComplexity int) int
	}
}

type MutationResolver interface {
	InitiateTest(ctx context.Context) (string, error)
	DestroyAll(ctx context.Context) (string, error)
	StartTest(ctx context.Context, input model.StartTestInput) (string, error)
	SaveTestResult(ctx context.Context, input model.SaveTestResultInput) (*model.Test, error)
	UpdateTest(ctx context.Context, input model.UpdateTestInput) (*model.Test, error)
	SeleniumPull(ctx context.Context) (string, error)
	SeleniumStart(ctx context.Context, input model.StartTestInput) (string, error)
	SeleniumStop(ctx context.Context, input model.SeleniumInput) (string, error)
	SeleniumRemove(ctx context.Context, input model.SeleniumInput) (string, error)
	CreateDockerFile(ctx context.Context, input model.CreateDockerFileInput) (string, error)
	CloneGitRepository(ctx context.Context, input model.CloneGitRepInput) (string, error)
	ProjectBuildImage(ctx context.Context, input model.StartTestInput) (string, error)
	ProjectContainerStart(ctx context.Context, input model.StartTestInput) (string, error)
	ProjectContainerStop(ctx context.Context, input model.ProjectContainerEndInput) (string, error)
	ProjectContainerRemove(ctx context.Context, input model.ProjectContainerEndInput) (string, error)
}
type QueryResolver interface {
	Tests(ctx context.Context) ([]*model.Test, error)
	Test(ctx context.Context, id int) (*model.Test, error)
	Settings(ctx context.Context) ([]*model.Settings, error)
	Setting(ctx context.Context, id int) (*model.Settings, error)
	ProjectInfos(ctx context.Context) ([]*model.Projectinfos, error)
	ProjectInfo(ctx context.Context, id int) (*model.Projectinfos, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.cloneGitRepository":
		if e.complexity.Mutation.CloneGitRepository == nil {
			break
		}

		args, err := ec.field_Mutation_cloneGitRepository_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CloneGitRepository(childComplexity, args["input"].(model.CloneGitRepInput)), true

	case "Mutation.createDockerFile":
		if e.complexity.Mutation.CreateDockerFile == nil {
			break
		}

		args, err := ec.field_Mutation_createDockerFile_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateDockerFile(childComplexity, args["input"].(model.CreateDockerFileInput)), true

	case "Mutation.destroyAll":
		if e.complexity.Mutation.DestroyAll == nil {
			break
		}

		return e.complexity.Mutation.DestroyAll(childComplexity), true

	case "Mutation.initiateTest":
		if e.complexity.Mutation.InitiateTest == nil {
			break
		}

		return e.complexity.Mutation.InitiateTest(childComplexity), true

	case "Mutation.projectBuildImage":
		if e.complexity.Mutation.ProjectBuildImage == nil {
			break
		}

		args, err := ec.field_Mutation_projectBuildImage_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ProjectBuildImage(childComplexity, args["input"].(model.StartTestInput)), true

	case "Mutation.projectContainerRemove":
		if e.complexity.Mutation.ProjectContainerRemove == nil {
			break
		}

		args, err := ec.field_Mutation_projectContainerRemove_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ProjectContainerRemove(childComplexity, args["input"].(model.ProjectContainerEndInput)), true

	case "Mutation.projectContainerStart":
		if e.complexity.Mutation.ProjectContainerStart == nil {
			break
		}

		args, err := ec.field_Mutation_projectContainerStart_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ProjectContainerStart(childComplexity, args["input"].(model.StartTestInput)), true

	case "Mutation.projectContainerStop":
		if e.complexity.Mutation.ProjectContainerStop == nil {
			break
		}

		args, err := ec.field_Mutation_projectContainerStop_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ProjectContainerStop(childComplexity, args["input"].(model.ProjectContainerEndInput)), true

	case "Mutation.saveTestResult":
		if e.complexity.Mutation.SaveTestResult == nil {
			break
		}

		args, err := ec.field_Mutation_saveTestResult_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SaveTestResult(childComplexity, args["input"].(model.SaveTestResultInput)), true

	case "Mutation.seleniumPull":
		if e.complexity.Mutation.SeleniumPull == nil {
			break
		}

		return e.complexity.Mutation.SeleniumPull(childComplexity), true

	case "Mutation.seleniumRemove":
		if e.complexity.Mutation.SeleniumRemove == nil {
			break
		}

		args, err := ec.field_Mutation_seleniumRemove_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SeleniumRemove(childComplexity, args["input"].(model.SeleniumInput)), true

	case "Mutation.seleniumStart":
		if e.complexity.Mutation.SeleniumStart == nil {
			break
		}

		args, err := ec.field_Mutation_seleniumStart_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SeleniumStart(childComplexity, args["input"].(model.StartTestInput)), true

	case "Mutation.seleniumStop":
		if e.complexity.Mutation.SeleniumStop == nil {
			break
		}

		args, err := ec.field_Mutation_seleniumStop_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SeleniumStop(childComplexity, args["input"].(model.SeleniumInput)), true

	case "Mutation.startTest":
		if e.complexity.Mutation.StartTest == nil {
			break
		}

		args, err := ec.field_Mutation_startTest_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StartTest(childComplexity, args["input"].(model.StartTestInput)), true

	case "Mutation.updateTest":
		if e.complexity.Mutation.UpdateTest == nil {
			break
		}

		args, err := ec.field_Mutation_updateTest_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateTest(childComplexity, args["input"].(model.UpdateTestInput)), true

	case "Projectinfos.gitEmail":
		if e.complexity.Projectinfos.GitEmail == nil {
			break
		}

		return e.complexity.Projectinfos.GitEmail(childComplexity), true

	case "Projectinfos.gitLanguage":
		if e.complexity.Projectinfos.GitLanguage == nil {
			break
		}

		return e.complexity.Projectinfos.GitLanguage(childComplexity), true

	case "Projectinfos.gitNoOfTestCases":
		if e.complexity.Projectinfos.GitNoOfTestCases == nil {
			break
		}

		return e.complexity.Projectinfos.GitNoOfTestCases(childComplexity), true

	case "Projectinfos.gitProjectName":
		if e.complexity.Projectinfos.GitProjectName == nil {
			break
		}

		return e.complexity.Projectinfos.GitProjectName(childComplexity), true

	case "Projectinfos.gitTestCaseFileName":
		if e.complexity.Projectinfos.GitTestCaseFileName == nil {
			break
		}

		return e.complexity.Projectinfos.GitTestCaseFileName(childComplexity), true

	case "Projectinfos.id":
		if e.complexity.Projectinfos.ID == nil {
			break
		}

		return e.complexity.Projectinfos.ID(childComplexity), true

	case "Query.projectInfo":
		if e.complexity.Query.ProjectInfo == nil {
			break
		}

		args, err := ec.field_Query_projectInfo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.ProjectInfo(childComplexity, args["id"].(int)), true

	case "Query.projectInfos":
		if e.complexity.Query.ProjectInfos == nil {
			break
		}

		return e.complexity.Query.ProjectInfos(childComplexity), true

	case "Query.setting":
		if e.complexity.Query.Setting == nil {
			break
		}

		args, err := ec.field_Query_setting_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Setting(childComplexity, args["id"].(int)), true

	case "Query.settings":
		if e.complexity.Query.Settings == nil {
			break
		}

		return e.complexity.Query.Settings(childComplexity), true

	case "Query.test":
		if e.complexity.Query.Test == nil {
			break
		}

		args, err := ec.field_Query_test_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Test(childComplexity, args["id"].(int)), true

	case "Query.tests":
		if e.complexity.Query.Tests == nil {
			break
		}

		return e.complexity.Query.Tests(childComplexity), true

	case "Settings.browser":
		if e.complexity.Settings.Browser == nil {
			break
		}

		return e.complexity.Settings.Browser(childComplexity), true

	case "Settings.enableLogs":
		if e.complexity.Settings.EnableLogs == nil {
			break
		}

		return e.complexity.Settings.EnableLogs(childComplexity), true

	case "Settings.id":
		if e.complexity.Settings.ID == nil {
			break
		}

		return e.complexity.Settings.ID(childComplexity), true

	case "Settings.numberOfParallelTests":
		if e.complexity.Settings.NumberOfParallelTests == nil {
			break
		}

		return e.complexity.Settings.NumberOfParallelTests(childComplexity), true

	case "Settings.parallelism":
		if e.complexity.Settings.Parallelism == nil {
			break
		}

		return e.complexity.Settings.Parallelism(childComplexity), true

	case "Settings.stepByStepDebugging":
		if e.complexity.Settings.StepByStepDebugging == nil {
			break
		}

		return e.complexity.Settings.StepByStepDebugging(childComplexity), true

	case "Settings.version":
		if e.complexity.Settings.Version == nil {
			break
		}

		return e.complexity.Settings.Version(childComplexity), true

	case "Test.duration":
		if e.complexity.Test.Duration == nil {
			break
		}

		return e.complexity.Test.Duration(childComplexity), true

	case "Test.id":
		if e.complexity.Test.ID == nil {
			break
		}

		return e.complexity.Test.ID(childComplexity), true

	case "Test.project":
		if e.complexity.Test.Project == nil {
			break
		}

		return e.complexity.Test.Project(childComplexity), true

	case "Test.status":
		if e.complexity.Test.Status == nil {
			break
		}

		return e.complexity.Test.Status(childComplexity), true

	case "Test.testfile":
		if e.complexity.Test.Testfile == nil {
			break
		}

		return e.complexity.Test.Testfile(childComplexity), true

	case "Test.type":
		if e.complexity.Test.Type == nil {
			break
		}

		return e.complexity.Test.Type(childComplexity), true

	case "Test.urlid":
		if e.complexity.Test.Urlid == nil {
			break
		}

		return e.complexity.Test.Urlid(childComplexity), true

	case "Test.username":
		if e.complexity.Test.Username == nil {
			break
		}

		return e.complexity.Test.Username(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputcloneGitRepInput,
		ec.unmarshalInputcreateDockerFileInput,
		ec.unmarshalInputprojectBuildImageInput,
		ec.unmarshalInputprojectContainerEndInput,
		ec.unmarshalInputprojectContainerStartInput,
		ec.unmarshalInputsaveTestResultInput,
		ec.unmarshalInputseleniumInput,
		ec.unmarshalInputstartTestInput,
		ec.unmarshalInputupdateTestInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

//go:embed "schema.graphqls"
var sourcesFS embed.FS

func sourceData(filename string) string {
	data, err := sourcesFS.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("codegen problem: %s not available", filename))
	}
	return string(data)
}

var sources = []*ast.Source{
	{Name: "schema.graphqls", Input: sourceData("schema.graphqls"), BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_cloneGitRepository_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.CloneGitRepInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNcloneGitRepInput2test_executorᚋgraphᚋmodelᚐCloneGitRepInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_createDockerFile_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.CreateDockerFileInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNcreateDockerFileInput2test_executorᚋgraphᚋmodelᚐCreateDockerFileInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_projectBuildImage_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.StartTestInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNstartTestInput2test_executorᚋgraphᚋmodelᚐStartTestInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_projectContainerRemove_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.ProjectContainerEndInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNprojectContainerEndInput2test_executorᚋgraphᚋmodelᚐProjectContainerEndInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_projectContainerStart_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.StartTestInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNstartTestInput2test_executorᚋgraphᚋmodelᚐStartTestInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_projectContainerStop_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.ProjectContainerEndInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNprojectContainerEndInput2test_executorᚋgraphᚋmodelᚐProjectContainerEndInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_saveTestResult_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.SaveTestResultInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNsaveTestResultInput2test_executorᚋgraphᚋmodelᚐSaveTestResultInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_seleniumRemove_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.SeleniumInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNseleniumInput2test_executorᚋgraphᚋmodelᚐSeleniumInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_seleniumStart_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.StartTestInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNstartTestInput2test_executorᚋgraphᚋmodelᚐStartTestInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_seleniumStop_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.SeleniumInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNseleniumInput2test_executorᚋgraphᚋmodelᚐSeleniumInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_startTest_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.StartTestInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNstartTestInput2test_executorᚋgraphᚋmodelᚐStartTestInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_updateTest_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.UpdateTestInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNupdateTestInput2test_executorᚋgraphᚋmodelᚐUpdateTestInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_projectInfo_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_setting_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_test_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("includeDeprecated"))
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("includeDeprecated"))
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_initiateTest(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_initiateTest(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().InitiateTest(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_initiateTest(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_destroyAll(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_destroyAll(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DestroyAll(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_destroyAll(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_startTest(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_startTest(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().StartTest(rctx, fc.Args["input"].(model.StartTestInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_startTest(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_startTest_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_saveTestResult(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_saveTestResult(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SaveTestResult(rctx, fc.Args["input"].(model.SaveTestResultInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Test)
	fc.Result = res
	return ec.marshalNTest2ᚖtest_executorᚋgraphᚋmodelᚐTest(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_saveTestResult(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Test_id(ctx, field)
			case "username":
				return ec.fieldContext_Test_username(ctx, field)
			case "project":
				return ec.fieldContext_Test_project(ctx, field)
			case "type":
				return ec.fieldContext_Test_type(ctx, field)
			case "urlid":
				return ec.fieldContext_Test_urlid(ctx, field)
			case "testfile":
				return ec.fieldContext_Test_testfile(ctx, field)
			case "status":
				return ec.fieldContext_Test_status(ctx, field)
			case "duration":
				return ec.fieldContext_Test_duration(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Test", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_saveTestResult_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateTest(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateTest(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateTest(rctx, fc.Args["input"].(model.UpdateTestInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Test)
	fc.Result = res
	return ec.marshalNTest2ᚖtest_executorᚋgraphᚋmodelᚐTest(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateTest(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Test_id(ctx, field)
			case "username":
				return ec.fieldContext_Test_username(ctx, field)
			case "project":
				return ec.fieldContext_Test_project(ctx, field)
			case "type":
				return ec.fieldContext_Test_type(ctx, field)
			case "urlid":
				return ec.fieldContext_Test_urlid(ctx, field)
			case "testfile":
				return ec.fieldContext_Test_testfile(ctx, field)
			case "status":
				return ec.fieldContext_Test_status(ctx, field)
			case "duration":
				return ec.fieldContext_Test_duration(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Test", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateTest_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_seleniumPull(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_seleniumPull(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SeleniumPull(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_seleniumPull(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_seleniumStart(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_seleniumStart(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SeleniumStart(rctx, fc.Args["input"].(model.StartTestInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_seleniumStart(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_seleniumStart_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_seleniumStop(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_seleniumStop(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SeleniumStop(rctx, fc.Args["input"].(model.SeleniumInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_seleniumStop(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_seleniumStop_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_seleniumRemove(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_seleniumRemove(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SeleniumRemove(rctx, fc.Args["input"].(model.SeleniumInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_seleniumRemove(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_seleniumRemove_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_createDockerFile(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createDockerFile(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateDockerFile(rctx, fc.Args["input"].(model.CreateDockerFileInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createDockerFile(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createDockerFile_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_cloneGitRepository(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_cloneGitRepository(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CloneGitRepository(rctx, fc.Args["input"].(model.CloneGitRepInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_cloneGitRepository(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_cloneGitRepository_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_projectBuildImage(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_projectBuildImage(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().ProjectBuildImage(rctx, fc.Args["input"].(model.StartTestInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_projectBuildImage(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_projectBuildImage_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_projectContainerStart(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_projectContainerStart(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().ProjectContainerStart(rctx, fc.Args["input"].(model.StartTestInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_projectContainerStart(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_projectContainerStart_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_projectContainerStop(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_projectContainerStop(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().ProjectContainerStop(rctx, fc.Args["input"].(model.ProjectContainerEndInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_projectContainerStop(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_projectContainerStop_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_projectContainerRemove(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_projectContainerRemove(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().ProjectContainerRemove(rctx, fc.Args["input"].(model.ProjectContainerEndInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_projectContainerRemove(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_projectContainerRemove_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Projectinfos_id(ctx context.Context, field graphql.CollectedField, obj *model.Projectinfos) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Projectinfos_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Projectinfos_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Projectinfos",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Projectinfos_gitEmail(ctx context.Context, field graphql.CollectedField, obj *model.Projectinfos) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Projectinfos_gitEmail(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.GitEmail, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Projectinfos_gitEmail(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Projectinfos",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Projectinfos_gitProjectName(ctx context.Context, field graphql.CollectedField, obj *model.Projectinfos) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Projectinfos_gitProjectName(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.GitProjectName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Projectinfos_gitProjectName(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Projectinfos",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Projectinfos_gitLanguage(ctx context.Context, field graphql.CollectedField, obj *model.Projectinfos) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Projectinfos_gitLanguage(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.GitLanguage, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Projectinfos_gitLanguage(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Projectinfos",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Projectinfos_gitNoOfTestCases(ctx context.Context, field graphql.CollectedField, obj *model.Projectinfos) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Projectinfos_gitNoOfTestCases(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.GitNoOfTestCases, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Projectinfos_gitNoOfTestCases(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Projectinfos",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Projectinfos_gitTestCaseFileName(ctx context.Context, field graphql.CollectedField, obj *model.Projectinfos) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Projectinfos_gitTestCaseFileName(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.GitTestCaseFileName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Projectinfos_gitTestCaseFileName(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Projectinfos",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query_tests(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_tests(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Tests(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Test)
	fc.Result = res
	return ec.marshalNTest2ᚕᚖtest_executorᚋgraphᚋmodelᚐTestᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_tests(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Test_id(ctx, field)
			case "username":
				return ec.fieldContext_Test_username(ctx, field)
			case "project":
				return ec.fieldContext_Test_project(ctx, field)
			case "type":
				return ec.fieldContext_Test_type(ctx, field)
			case "urlid":
				return ec.fieldContext_Test_urlid(ctx, field)
			case "testfile":
				return ec.fieldContext_Test_testfile(ctx, field)
			case "status":
				return ec.fieldContext_Test_status(ctx, field)
			case "duration":
				return ec.fieldContext_Test_duration(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Test", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query_test(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_test(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Test(rctx, fc.Args["id"].(int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.Test)
	fc.Result = res
	return ec.marshalOTest2ᚖtest_executorᚋgraphᚋmodelᚐTest(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_test(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Test_id(ctx, field)
			case "username":
				return ec.fieldContext_Test_username(ctx, field)
			case "project":
				return ec.fieldContext_Test_project(ctx, field)
			case "type":
				return ec.fieldContext_Test_type(ctx, field)
			case "urlid":
				return ec.fieldContext_Test_urlid(ctx, field)
			case "testfile":
				return ec.fieldContext_Test_testfile(ctx, field)
			case "status":
				return ec.fieldContext_Test_status(ctx, field)
			case "duration":
				return ec.fieldContext_Test_duration(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Test", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_test_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query_settings(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_settings(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Settings(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Settings)
	fc.Result = res
	return ec.marshalNSettings2ᚕᚖtest_executorᚋgraphᚋmodelᚐSettingsᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_settings(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Settings_id(ctx, field)
			case "browser":
				return ec.fieldContext_Settings_browser(ctx, field)
			case "version":
				return ec.fieldContext_Settings_version(ctx, field)
			case "stepByStepDebugging":
				return ec.fieldContext_Settings_stepByStepDebugging(ctx, field)
			case "enableLogs":
				return ec.fieldContext_Settings_enableLogs(ctx, field)
			case "parallelism":
				return ec.fieldContext_Settings_parallelism(ctx, field)
			case "numberOfParallelTests":
				return ec.fieldContext_Settings_numberOfParallelTests(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Settings", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query_setting(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_setting(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Setting(rctx, fc.Args["id"].(int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Settings)
	fc.Result = res
	return ec.marshalNSettings2ᚖtest_executorᚋgraphᚋmodelᚐSettings(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_setting(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Settings_id(ctx, field)
			case "browser":
				return ec.fieldContext_Settings_browser(ctx, field)
			case "version":
				return ec.fieldContext_Settings_version(ctx, field)
			case "stepByStepDebugging":
				return ec.fieldContext_Settings_stepByStepDebugging(ctx, field)
			case "enableLogs":
				return ec.fieldContext_Settings_enableLogs(ctx, field)
			case "parallelism":
				return ec.fieldContext_Settings_parallelism(ctx, field)
			case "numberOfParallelTests":
				return ec.fieldContext_Settings_numberOfParallelTests(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Settings", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_setting_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query_projectInfos(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_projectInfos(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ProjectInfos(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Projectinfos)
	fc.Result = res
	return ec.marshalNProjectinfos2ᚕᚖtest_executorᚋgraphᚋmodelᚐProjectinfosᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_projectInfos(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Projectinfos_id(ctx, field)
			case "gitEmail":
				return ec.fieldContext_Projectinfos_gitEmail(ctx, field)
			case "gitProjectName":
				return ec.fieldContext_Projectinfos_gitProjectName(ctx, field)
			case "gitLanguage":
				return ec.fieldContext_Projectinfos_gitLanguage(ctx, field)
			case "gitNoOfTestCases":
				return ec.fieldContext_Projectinfos_gitNoOfTestCases(ctx, field)
			case "gitTestCaseFileName":
				return ec.fieldContext_Projectinfos_gitTestCaseFileName(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Projectinfos", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query_projectInfo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_projectInfo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ProjectInfo(rctx, fc.Args["id"].(int))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Projectinfos)
	fc.Result = res
	return ec.marshalNProjectinfos2ᚖtest_executorᚋgraphᚋmodelᚐProjectinfos(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_projectInfo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Projectinfos_id(ctx, field)
			case "gitEmail":
				return ec.fieldContext_Projectinfos_gitEmail(ctx, field)
			case "gitProjectName":
				return ec.fieldContext_Projectinfos_gitProjectName(ctx, field)
			case "gitLanguage":
				return ec.fieldContext_Projectinfos_gitLanguage(ctx, field)
			case "gitNoOfTestCases":
				return ec.fieldContext_Projectinfos_gitNoOfTestCases(ctx, field)
			case "gitTestCaseFileName":
				return ec.fieldContext_Projectinfos_gitTestCaseFileName(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Projectinfos", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_projectInfo_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(fc.Args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query___type_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___schema(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___schema(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "description":
				return ec.fieldContext___Schema_description(ctx, field)
			case "types":
				return ec.fieldContext___Schema_types(ctx, field)
			case "queryType":
				return ec.fieldContext___Schema_queryType(ctx, field)
			case "mutationType":
				return ec.fieldContext___Schema_mutationType(ctx, field)
			case "subscriptionType":
				return ec.fieldContext___Schema_subscriptionType(ctx, field)
			case "directives":
				return ec.fieldContext___Schema_directives(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Schema", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Settings_id(ctx context.Context, field graphql.CollectedField, obj *model.Settings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Settings_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Settings_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Settings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Settings_browser(ctx context.Context, field graphql.CollectedField, obj *model.Settings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Settings_browser(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Browser, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Settings_browser(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Settings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Settings_version(ctx context.Context, field graphql.CollectedField, obj *model.Settings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Settings_version(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Version, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Settings_version(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Settings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Settings_stepByStepDebugging(ctx context.Context, field graphql.CollectedField, obj *model.Settings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Settings_stepByStepDebugging(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.StepByStepDebugging, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Settings_stepByStepDebugging(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Settings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Settings_enableLogs(ctx context.Context, field graphql.CollectedField, obj *model.Settings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Settings_enableLogs(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnableLogs, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Settings_enableLogs(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Settings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Settings_parallelism(ctx context.Context, field graphql.CollectedField, obj *model.Settings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Settings_parallelism(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Parallelism, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Settings_parallelism(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Settings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Settings_numberOfParallelTests(ctx context.Context, field graphql.CollectedField, obj *model.Settings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Settings_numberOfParallelTests(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NumberOfParallelTests, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Settings_numberOfParallelTests(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Settings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_id(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_username(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_username(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Username, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_username(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_project(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_project(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Project, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_project(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_type(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_urlid(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_urlid(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Urlid, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_urlid(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_testfile(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_testfile(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Testfile, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_testfile(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_status(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_status(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Status, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_status(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Test_duration(ctx context.Context, field graphql.CollectedField, obj *model.Test) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Test_duration(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Duration, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Test_duration(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Test",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Directive_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Directive_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Directive_description(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Directive_description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Directive_locations(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalN__DirectiveLocation2ᚕstringᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Directive_locations(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type __DirectiveLocation does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Directive_args(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Directive_args(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext___InputValue_name(ctx, field)
			case "description":
				return ec.fieldContext___InputValue_description(ctx, field)
			case "type":
				return ec.fieldContext___InputValue_type(ctx, field)
			case "defaultValue":
				return ec.fieldContext___InputValue_defaultValue(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __InputValue", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Directive_isRepeatable(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Directive_isRepeatable(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsRepeatable, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Directive_isRepeatable(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Directive",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___EnumValue_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___EnumValue_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___EnumValue_description(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___EnumValue_description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___EnumValue_isDeprecated(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___EnumValue_deprecationReason(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__EnumValue",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Field_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Field_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Field_description(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Field_description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Field_args(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Field_args(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext___InputValue_name(ctx, field)
			case "description":
				return ec.fieldContext___InputValue_description(ctx, field)
			case "type":
				return ec.fieldContext___InputValue_type(ctx, field)
			case "defaultValue":
				return ec.fieldContext___InputValue_defaultValue(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __InputValue", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Field_type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Field_type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Field_isDeprecated(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Field_isDeprecated(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Field_deprecationReason(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Field_deprecationReason(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Field",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___InputValue_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___InputValue_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___InputValue_description(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___InputValue_description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___InputValue_type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___InputValue_type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___InputValue_defaultValue(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__InputValue",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Schema_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Schema_description(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Schema_description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Schema_types(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Schema_types(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Schema_queryType(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Schema_queryType(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Schema_mutationType(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Schema_mutationType(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Schema_subscriptionType(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Schema_directives(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	fc.Result = res
	return ec.marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirectiveᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Schema_directives(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Schema",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext___Directive_name(ctx, field)
			case "description":
				return ec.fieldContext___Directive_description(ctx, field)
			case "locations":
				return ec.fieldContext___Directive_locations(ctx, field)
			case "args":
				return ec.fieldContext___Directive_args(ctx, field)
			case "isRepeatable":
				return ec.fieldContext___Directive_isRepeatable(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Directive", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_kind(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalN__TypeKind2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_kind(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type __TypeKind does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_description(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_fields(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(fc.Args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	fc.Result = res
	return ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐFieldᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_fields(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext___Field_name(ctx, field)
			case "description":
				return ec.fieldContext___Field_description(ctx, field)
			case "args":
				return ec.fieldContext___Field_args(ctx, field)
			case "type":
				return ec.fieldContext___Field_type(ctx, field)
			case "isDeprecated":
				return ec.fieldContext___Field_isDeprecated(ctx, field)
			case "deprecationReason":
				return ec.fieldContext___Field_deprecationReason(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Field", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field___Type_fields_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_interfaces(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_interfaces(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_possibleTypes(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_possibleTypes(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_enumValues(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(fc.Args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	fc.Result = res
	return ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_enumValues(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext___EnumValue_name(ctx, field)
			case "description":
				return ec.fieldContext___EnumValue_description(ctx, field)
			case "isDeprecated":
				return ec.fieldContext___EnumValue_isDeprecated(ctx, field)
			case "deprecationReason":
				return ec.fieldContext___EnumValue_deprecationReason(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __EnumValue", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field___Type_enumValues_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_inputFields(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_inputFields(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext___InputValue_name(ctx, field)
			case "description":
				return ec.fieldContext___InputValue_description(ctx, field)
			case "type":
				return ec.fieldContext___InputValue_type(ctx, field)
			case "defaultValue":
				return ec.fieldContext___InputValue_defaultValue(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __InputValue", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_ofType(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_ofType(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) ___Type_specifiedByURL(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext___Type_specifiedByURL(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SpecifiedByURL(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext___Type_specifiedByURL(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "__Type",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputcloneGitRepInput(ctx context.Context, obj interface{}) (model.CloneGitRepInput, error) {
	var it model.CloneGitRepInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"gitRepPath"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "gitRepPath":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("gitRepPath"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.GitRepPath = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputcreateDockerFileInput(ctx context.Context, obj interface{}) (model.CreateDockerFileInput, error) {
	var it model.CreateDockerFileInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"dockerCmnds"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "dockerCmnds":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dockerCmnds"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.DockerCmnds = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputprojectBuildImageInput(ctx context.Context, obj interface{}) (model.ProjectBuildImageInput, error) {
	var it model.ProjectBuildImageInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"imgName", "projectPath"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "imgName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("imgName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ImgName = data
		case "projectPath":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("projectPath"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ProjectPath = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputprojectContainerEndInput(ctx context.Context, obj interface{}) (model.ProjectContainerEndInput, error) {
	var it model.ProjectContainerEndInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"containerName"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "containerName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("containerName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ContainerName = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputprojectContainerStartInput(ctx context.Context, obj interface{}) (model.ProjectContainerStartInput, error) {
	var it model.ProjectContainerStartInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"imageName", "containerName"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "imageName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("imageName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ImageName = data
		case "containerName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("containerName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ContainerName = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputsaveTestResultInput(ctx context.Context, obj interface{}) (model.SaveTestResultInput, error) {
	var it model.SaveTestResultInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"username", "project", "type", "urlid", "testfile", "status", "duration"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "username":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("username"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Username = data
		case "project":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("project"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Project = data
		case "type":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Type = data
		case "urlid":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("urlid"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Urlid = data
		case "testfile":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("testfile"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Testfile = data
		case "status":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("status"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Status = data
		case "duration":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("duration"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Duration = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputseleniumInput(ctx context.Context, obj interface{}) (model.SeleniumInput, error) {
	var it model.SeleniumInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"containerName"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "containerName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("containerName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ContainerName = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputstartTestInput(ctx context.Context, obj interface{}) (model.StartTestInput, error) {
	var it model.StartTestInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"seleniumContainerName", "projectPath", "projectImageName", "projectContainerName"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "seleniumContainerName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("seleniumContainerName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.SeleniumContainerName = data
		case "projectPath":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("projectPath"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ProjectPath = data
		case "projectImageName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("projectImageName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ProjectImageName = data
		case "projectContainerName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("projectContainerName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ProjectContainerName = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputupdateTestInput(ctx context.Context, obj interface{}) (model.UpdateTestInput, error) {
	var it model.UpdateTestInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "status", "testPath"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNID2int(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "status":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("status"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Status = data
		case "testPath":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("testPath"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.TestPath = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "initiateTest":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_initiateTest(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "destroyAll":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_destroyAll(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "startTest":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_startTest(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "saveTestResult":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_saveTestResult(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "updateTest":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateTest(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "seleniumPull":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_seleniumPull(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "seleniumStart":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_seleniumStart(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "seleniumStop":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_seleniumStop(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "seleniumRemove":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_seleniumRemove(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "createDockerFile":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createDockerFile(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "cloneGitRepository":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_cloneGitRepository(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "projectBuildImage":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_projectBuildImage(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "projectContainerStart":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_projectContainerStart(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "projectContainerStop":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_projectContainerStop(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "projectContainerRemove":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_projectContainerRemove(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var projectinfosImplementors = []string{"Projectinfos"}

func (ec *executionContext) _Projectinfos(ctx context.Context, sel ast.SelectionSet, obj *model.Projectinfos) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, projectinfosImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Projectinfos")
		case "id":

			out.Values[i] = ec._Projectinfos_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "gitEmail":

			out.Values[i] = ec._Projectinfos_gitEmail(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "gitProjectName":

			out.Values[i] = ec._Projectinfos_gitProjectName(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "gitLanguage":

			out.Values[i] = ec._Projectinfos_gitLanguage(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "gitNoOfTestCases":

			out.Values[i] = ec._Projectinfos_gitNoOfTestCases(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "gitTestCaseFileName":

			out.Values[i] = ec._Projectinfos_gitTestCaseFileName(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "tests":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_tests(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "test":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_test(ctx, field)
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "settings":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_settings(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "setting":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_setting(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "projectInfos":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_projectInfos(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "projectInfo":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_projectInfo(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "__type":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___type(ctx, field)
			})

		case "__schema":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___schema(ctx, field)
			})

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var settingsImplementors = []string{"Settings"}

func (ec *executionContext) _Settings(ctx context.Context, sel ast.SelectionSet, obj *model.Settings) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, settingsImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Settings")
		case "id":

			out.Values[i] = ec._Settings_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "browser":

			out.Values[i] = ec._Settings_browser(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "version":

			out.Values[i] = ec._Settings_version(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "stepByStepDebugging":

			out.Values[i] = ec._Settings_stepByStepDebugging(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "enableLogs":

			out.Values[i] = ec._Settings_enableLogs(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "parallelism":

			out.Values[i] = ec._Settings_parallelism(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "numberOfParallelTests":

			out.Values[i] = ec._Settings_numberOfParallelTests(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var testImplementors = []string{"Test"}

func (ec *executionContext) _Test(ctx context.Context, sel ast.SelectionSet, obj *model.Test) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, testImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Test")
		case "id":

			out.Values[i] = ec._Test_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "username":

			out.Values[i] = ec._Test_username(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "project":

			out.Values[i] = ec._Test_project(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":

			out.Values[i] = ec._Test_type(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "urlid":

			out.Values[i] = ec._Test_urlid(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "testfile":

			out.Values[i] = ec._Test_testfile(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "status":

			out.Values[i] = ec._Test_status(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "duration":

			out.Values[i] = ec._Test_duration(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __DirectiveImplementors = []string{"__Directive"}

func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __DirectiveImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":

			out.Values[i] = ec.___Directive_name(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":

			out.Values[i] = ec.___Directive_description(ctx, field, obj)

		case "locations":

			out.Values[i] = ec.___Directive_locations(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "args":

			out.Values[i] = ec.___Directive_args(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isRepeatable":

			out.Values[i] = ec.___Directive_isRepeatable(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __EnumValueImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":

			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":

			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)

		case "isDeprecated":

			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":

			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __FieldImplementors = []string{"__Field"}

func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __FieldImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":

			out.Values[i] = ec.___Field_name(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":

			out.Values[i] = ec.___Field_description(ctx, field, obj)

		case "args":

			out.Values[i] = ec.___Field_args(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":

			out.Values[i] = ec.___Field_type(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isDeprecated":

			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":

			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __InputValueImplementors = []string{"__InputValue"}

func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __InputValueImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":

			out.Values[i] = ec.___InputValue_name(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":

			out.Values[i] = ec.___InputValue_description(ctx, field, obj)

		case "type":

			out.Values[i] = ec.___InputValue_type(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "defaultValue":

			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __SchemaImplementors = []string{"__Schema"}

func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __SchemaImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "description":

			out.Values[i] = ec.___Schema_description(ctx, field, obj)

		case "types":

			out.Values[i] = ec.___Schema_types(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "queryType":

			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "mutationType":

			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)

		case "subscriptionType":

			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)

		case "directives":

			out.Values[i] = ec.___Schema_directives(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __TypeImplementors = []string{"__Type"}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __TypeImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":

			out.Values[i] = ec.___Type_kind(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":

			out.Values[i] = ec.___Type_name(ctx, field, obj)

		case "description":

			out.Values[i] = ec.___Type_description(ctx, field, obj)

		case "fields":

			out.Values[i] = ec.___Type_fields(ctx, field, obj)

		case "interfaces":

			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)

		case "possibleTypes":

			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)

		case "enumValues":

			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)

		case "inputFields":

			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)

		case "ofType":

			out.Values[i] = ec.___Type_ofType(ctx, field, obj)

		case "specifiedByURL":

			out.Values[i] = ec.___Type_specifiedByURL(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	res, err := graphql.UnmarshalBoolean(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	res := graphql.MarshalBoolean(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNID2int(ctx context.Context, v interface{}) (int, error) {
	res, err := graphql.UnmarshalInt(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNID2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	res := graphql.MarshalInt(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNInt2int(ctx context.Context, v interface{}) (int, error) {
	res, err := graphql.UnmarshalInt(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNInt2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	res := graphql.MarshalInt(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) marshalNProjectinfos2test_executorᚋgraphᚋmodelᚐProjectinfos(ctx context.Context, sel ast.SelectionSet, v model.Projectinfos) graphql.Marshaler {
	return ec._Projectinfos(ctx, sel, &v)
}

func (ec *executionContext) marshalNProjectinfos2ᚕᚖtest_executorᚋgraphᚋmodelᚐProjectinfosᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Projectinfos) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNProjectinfos2ᚖtest_executorᚋgraphᚋmodelᚐProjectinfos(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNProjectinfos2ᚖtest_executorᚋgraphᚋmodelᚐProjectinfos(ctx context.Context, sel ast.SelectionSet, v *model.Projectinfos) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Projectinfos(ctx, sel, v)
}

func (ec *executionContext) marshalNSettings2test_executorᚋgraphᚋmodelᚐSettings(ctx context.Context, sel ast.SelectionSet, v model.Settings) graphql.Marshaler {
	return ec._Settings(ctx, sel, &v)
}

func (ec *executionContext) marshalNSettings2ᚕᚖtest_executorᚋgraphᚋmodelᚐSettingsᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Settings) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNSettings2ᚖtest_executorᚋgraphᚋmodelᚐSettings(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNSettings2ᚖtest_executorᚋgraphᚋmodelᚐSettings(ctx context.Context, sel ast.SelectionSet, v *model.Settings) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Settings(ctx, sel, v)
}

func (ec *executionContext) unmarshalNString2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) marshalNTest2test_executorᚋgraphᚋmodelᚐTest(ctx context.Context, sel ast.SelectionSet, v model.Test) graphql.Marshaler {
	return ec._Test(ctx, sel, &v)
}

func (ec *executionContext) marshalNTest2ᚕᚖtest_executorᚋgraphᚋmodelᚐTestᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Test) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNTest2ᚖtest_executorᚋgraphᚋmodelᚐTest(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNTest2ᚖtest_executorᚋgraphᚋmodelᚐTest(ctx context.Context, sel ast.SelectionSet, v *model.Test) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Test(ctx, sel, v)
}

func (ec *executionContext) marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v introspection.Directive) graphql.Marshaler {
	return ec.___Directive(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirectiveᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Directive) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalN__DirectiveLocation2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalN__DirectiveLocation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalN__DirectiveLocation2ᚕstringᚄ(ctx context.Context, v interface{}) ([]string, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]string, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalN__DirectiveLocation2string(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalN__DirectiveLocation2ᚕstringᚄ(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__DirectiveLocation2string(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v introspection.EnumValue) graphql.Marshaler {
	return ec.___EnumValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v introspection.Field) graphql.Marshaler {
	return ec.___Field(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v introspection.InputValue) graphql.Marshaler {
	return ec.___InputValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

func (ec *executionContext) unmarshalN__TypeKind2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalN__TypeKind2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNcloneGitRepInput2test_executorᚋgraphᚋmodelᚐCloneGitRepInput(ctx context.Context, v interface{}) (model.CloneGitRepInput, error) {
	res, err := ec.unmarshalInputcloneGitRepInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNcreateDockerFileInput2test_executorᚋgraphᚋmodelᚐCreateDockerFileInput(ctx context.Context, v interface{}) (model.CreateDockerFileInput, error) {
	res, err := ec.unmarshalInputcreateDockerFileInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNprojectContainerEndInput2test_executorᚋgraphᚋmodelᚐProjectContainerEndInput(ctx context.Context, v interface{}) (model.ProjectContainerEndInput, error) {
	res, err := ec.unmarshalInputprojectContainerEndInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNsaveTestResultInput2test_executorᚋgraphᚋmodelᚐSaveTestResultInput(ctx context.Context, v interface{}) (model.SaveTestResultInput, error) {
	res, err := ec.unmarshalInputsaveTestResultInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNseleniumInput2test_executorᚋgraphᚋmodelᚐSeleniumInput(ctx context.Context, v interface{}) (model.SeleniumInput, error) {
	res, err := ec.unmarshalInputseleniumInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNstartTestInput2test_executorᚋgraphᚋmodelᚐStartTestInput(ctx context.Context, v interface{}) (model.StartTestInput, error) {
	res, err := ec.unmarshalInputstartTestInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNupdateTestInput2test_executorᚋgraphᚋmodelᚐUpdateTestInput(ctx context.Context, v interface{}) (model.UpdateTestInput, error) {
	res, err := ec.unmarshalInputupdateTestInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	res, err := graphql.UnmarshalBoolean(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	res := graphql.MarshalBoolean(v)
	return res
}

func (ec *executionContext) unmarshalOBoolean2ᚖbool(ctx context.Context, v interface{}) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalBoolean(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOBoolean2ᚖbool(ctx context.Context, sel ast.SelectionSet, v *bool) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalBoolean(*v)
	return res
}

func (ec *executionContext) unmarshalOString2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalString(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOString2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalString(*v)
	return res
}

func (ec *executionContext) marshalOTest2ᚖtest_executorᚋgraphᚋmodelᚐTest(ctx context.Context, sel ast.SelectionSet, v *model.Test) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Test(ctx, sel, v)
}

func (ec *executionContext) marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValueᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.EnumValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐFieldᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Field) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v *introspection.Schema) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, sel, v)
}

func (ec *executionContext) marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
