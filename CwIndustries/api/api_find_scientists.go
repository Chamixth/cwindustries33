package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
  "encoding/json"
)


// @Summary      GET Scientists input: Scientists
// @Description  GET Scientists API
// @Tags         GET Scientists - Scientists
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Scientists "Status OK"
// @Success      202  {array}   dto.Model_Scientists "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindScientists [GET]


    func FindScientistsApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


objectid := event.QueryStringParameters["objectid"]
    inputObj, err := dao.DB_FindScientists(objectid)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)

    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil

}
