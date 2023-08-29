package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  )


// @Summary      DELETE Projects input: Projects
// @Description  DELETE Projects API
// @Tags         DELETE Projects - Projects
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Projects "Status OK"
// @Success      202  {array}   dto.Model_Projects "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteProjects [DELETE]


    func DeleteProjectsApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


objectid := event.QueryStringParameters["objectid"]
    err := dao.DB_DeleteProjects(objectid)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       string(""),
    },nil
    

}
