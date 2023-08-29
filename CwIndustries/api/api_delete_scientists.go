package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  )


// @Summary      DELETE Scientists input: Scientists
// @Description  DELETE Scientists API
// @Tags         DELETE Scientists - Scientists
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Scientists "Status OK"
// @Success      202  {array}   dto.Model_Scientists "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteScientists [DELETE]


    func DeleteScientistsApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


objectid := event.QueryStringParameters["objectid"]
    err := dao.DB_DeleteScientists(objectid)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       string(""),
    },nil
    

}
