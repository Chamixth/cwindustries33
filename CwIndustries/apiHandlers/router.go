package apiHandlers

import (
	"CwIndustries/api"
	"CwIndustries/dbConfig"

	"context"


	"github.com/aws/aws-lambda-go/events"
)



func Router(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
	dbConfig.ConnectToMongoDB()
	method := request.QueryStringParameters["method"]
	switch method {
case "CreateEngineers":
return api.CreateEngineersApi(ctx, request)

case "UpdateEngineers":
return api.UpdateEngineersApi(ctx, request)

case "DeleteEngineers":
return api.DeleteEngineersApi(ctx, request)

case "FindEngineers":
return api.FindEngineersApi(ctx, request)

case "FindallEngineers":
return api.FindallEngineersApi(ctx, request)

case "CreateScientists":
return api.CreateScientistsApi(ctx, request)

case "UpdateScientists":
return api.UpdateScientistsApi(ctx, request)

case "DeleteScientists":
return api.DeleteScientistsApi(ctx, request)

case "FindScientists":
return api.FindScientistsApi(ctx, request)

case "FindallScientists":
return api.FindallScientistsApi(ctx, request)

case "CreateProjects":
return api.CreateProjectsApi(ctx, request)

case "UpdateProjects":
return api.UpdateProjectsApi(ctx, request)

case "DeleteProjects":
return api.DeleteProjectsApi(ctx, request)

case "FindProjects":
return api.FindProjectsApi(ctx, request)

case "FindallProjects":
return api.FindallProjectsApi(ctx, request)

case "CreateWeapons":
return api.CreateWeaponsApi(ctx, request)

case "UpdateWeapons":
return api.UpdateWeaponsApi(ctx, request)

case "DeleteWeapons":
return api.DeleteWeaponsApi(ctx, request)

case "FindWeapons":
return api.FindWeaponsApi(ctx, request)

case "FindallWeapons":
return api.FindallWeaponsApi(ctx, request)

default:
return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Method Not Allowed"},nil
}

}