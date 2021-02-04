// Guide (I am using bash and please make that two Docker build and run commands inside this file)\
// #################################################################################################

// @ GET request\
// Simply get the array back\
// curl -X GET http://localhost:80/list 

// @ POST request\
// 😀 = something you want\
// Set the body data type in header and the request body itself first \
// curl -X POST -H "Content-Type: application/json" -d '{"key":"😀", "value":"😀"}' http://localhost:80/add
