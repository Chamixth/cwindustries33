package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  )


// @Summary      DELETE Weapons input: Weapons
// @Description  DELETE Weapons API
// @Tags         DELETE Weapons - Weapons
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Weapons "Status OK"
// @Success      202  {array}   dto.Model_Weapons "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteWeapons [DELETE]


    func DeleteWeaponsApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


objectid := event.QueryStringParameters["objectid"]
    err := dao.DB_DeleteWeapons(objectid)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       string(""),
    },nil
    

}
