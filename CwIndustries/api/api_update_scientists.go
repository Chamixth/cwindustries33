package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
    "encoding/json"
    "CwIndustries/dto"

  )


// @Summary      PUT Scientists input: Scientists
// @Description  PUT Scientists API
// @Tags         PUT Scientists - Scientists
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Scientists false "string collection" 
// @Success      200  {array}   dto.Model_Scientists "Status OK"
// @Success      202  {array}   dto.Model_Scientists "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateScientists [PUT]


    func UpdateScientistsApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


inputObj := dto.Scientists{}
    if err := json.Unmarshal([]byte(event.Body), &inputObj); err != nil {
    		return  events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid Request Body"},nil
    }

    //validate := validator.New()
    //if validationErr := validate.Struct(&inputObj); validationErr != nil {
      //  return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    //}
err := dao.DB_UpdateScientists(&inputObj)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)
    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil
    
    

}
