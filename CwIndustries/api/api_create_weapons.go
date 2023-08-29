package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
    "encoding/json"
    "CwIndustries/dto"

  
    "CwIndustries/functions"
  

  )


// @Summary      POST Weapons input: Weapons
// @Description  POST Weapons API
// @Tags         POST Weapons - Weapons
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Weapons false "string collection" 
// @Success      200  {array}   dto.Model_Weapons "Status OK"
// @Success      202  {array}   dto.Model_Weapons "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateWeapons [POST]


    func CreateWeaponsApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


inputObj := dto.Weapons{}
    if err := json.Unmarshal([]byte(event.Body), &inputObj); err != nil {
    		return  events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid Request Body"},nil
    }

    //validate := validator.New()
    //if validationErr := validate.Struct(&inputObj); validationErr != nil {
      //  return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    //}
err := dao.DB_CreateWeapons(&inputObj)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)
    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil
    
    

}
