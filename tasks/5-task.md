- Open a file, if file does not exist , create the file
- Write Person struct data to the file in Json format
{"Name":"Jiten","Age":38,"Email":"JitenP@outlook.Com"} , use fmt.Springf
Give command line arguments like name, age and email 
Every time the binary is called along with command line arguments that should be saved in the file

-- go run main.go jiten 38 jitenp@outlook.com 
-- go run main.go keerthana 32 kn@gmail.com
-- go run main.go logeswern 34 lg@gmail.lcom
