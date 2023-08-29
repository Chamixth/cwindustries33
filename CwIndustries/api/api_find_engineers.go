package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
  "encoding/json"
)


// @Summary      GET Engineers input: Engineers
// @Description  GET Engineers API
// @Tags         GET Engineers - Engineers
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Engineers "Status OK"
// @Success      202  {array}   dto.Model_Engineers "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindEngineers [GET]


    func FindEngineersApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


objectid := event.QueryStringParameters["objectid"]
    inputObj, err := dao.DB_FindEngineers(objectid)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)

    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil

}
