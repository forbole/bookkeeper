package flow

import (
	"fmt"
	"io/ioutil"

	"github.com/forbole/bookkeeper/module/flow/utils"
	"github.com/forbole/bookkeeper/module/flow/tables"

	"github.com/forbole/bookkeeper/types"
)

func HandleNodeInfos(flow types.Flow)([]string,error){
	if len(flow.NodeIds)==0{
		return nil,nil
	}
	var filenames []string
	flowClient,err:=utils.NewFlowClient(flow.FlowEndpoint)
	if err!=nil{
		return nil,err
	}

	for _,id:=range flow.NodeIds{
		nodeInfo,err:=tables.GetNodeInfo(id,flow.FlowJuno)
		if err!=nil{
			return nil,err
		}

		outputcsv,err := nodeInfo.GetCSV(flow.Exponent,"flow","USD",*flowClient)
		if err!=nil{
			return nil,err
		}
		fmt.Println(outputcsv)
		filename := fmt.Sprintf("%s_nodeInfo.csv", id)
		err = ioutil.WriteFile(filename, []byte(outputcsv), 0777)
		if err != nil {
			return nil,err
		}
		filenames = append(filenames, filename)
	}
	return filenames,nil
	
}
