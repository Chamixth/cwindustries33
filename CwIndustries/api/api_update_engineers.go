package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
    "encoding/json"
    "CwIndustries/dto"

  )


// @Summary      PUT Engineers input: Engineers
// @Description  PUT Engineers API
// @Tags         PUT Engineers - Engineers
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Engineers false "string collection" 
// @Success      200  {array}   dto.Model_Engineers "Status OK"
// @Success      202  {array}   dto.Model_Engineers "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateEngineers [PUT]


    func UpdateEngineersApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


inputObj := dto.Engineers{}
    if err := json.Unmarshal([]byte(event.Body), &inputObj); err != nil {
    		return  events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid Request Body"},nil
    }

    //validate := validator.New()
    //if validationErr := validate.Struct(&inputObj); validationErr != nil {
      //  return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    //}
err := dao.DB_UpdateEngineers(&inputObj)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)
    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil
    
    

}
