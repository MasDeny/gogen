package gogen

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type datasource struct {
}

func NewDatasource() Generator {
	return &datasource{}
}

func (d *datasource) Generate(args ...string) error {

	// if IsNotExist(".application_schema/") {
	// 	return fmt.Errorf("please call `gogen init` first")
	// }

	if len(args) < 4 {
		return fmt.Errorf("please define datasource and usecase_name. ex: `gogen datasource production CreateOrder`")
	}

	datasourceName := args[2]

	usecaseName := args[3]

	ds := Datasource{}
	ds.DatasourceName = datasourceName
	ds.UsecaseName = usecaseName
	ds.PackagePath = GetPackagePath()

	{
		file, err := os.Open(fmt.Sprintf("usecase/%s/port/outport.go", strings.ToLower(usecaseName)))
		if err != nil {
			return fmt.Errorf("not found usecase %s. You need to create it first by call 'gogen usecase %s' ", usecaseName, usecaseName)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		state := 0
		for scanner.Scan() {
			if state == 0 && strings.HasPrefix(scanner.Text(), fmt.Sprintf("type %sOutport interface {", usecaseName)) {
				state = 1
			} else //
			if state == 1 {
				if strings.HasPrefix(scanner.Text(), "}") {
					state = 2
					break
				} else {
					completeMethod := strings.TrimSpace(scanner.Text())
					methodNameOnly := strings.Split(completeMethod, "(")[0]
					ds.Outports = append(ds.Outports, &Outport{
						Name: methodNameOnly,
					})
				}
			}
		}

		if state == 0 {
			return fmt.Errorf("not found usecase %s. You need to create it first by call 'gogen usecase %s' ", usecaseName, usecaseName)
		}
	}

	for _, ot := range ds.Outports {

		file, err := os.Open(fmt.Sprintf("usecase/%s/port/outport.go", strings.ToLower(usecaseName)))
		if err != nil {
			return fmt.Errorf("not found usecase %s. You need to create it first by call 'gogen usecase %s' ", usecaseName, usecaseName)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		state := 0
		for scanner.Scan() {
			if state == 0 && strings.HasPrefix(scanner.Text(), fmt.Sprintf("type %sRequest struct {", ot.Name)) {
				state = 1
			} else //
			if state == 1 {
				if strings.HasPrefix(scanner.Text(), "}") {
					state = 2
					break
				} else {

					completeFieldWithType := strings.TrimSpace(scanner.Text())
					if len(completeFieldWithType) == 0 {
						continue
					}
					fieldWithType := strings.SplitN(completeFieldWithType, " ", 2)
					ot.RequestFields = append(ot.RequestFields, &NameType{
						Name: strings.TrimSpace(fieldWithType[0]),
					})

				}
			}
		}

		if state == 0 {
			return fmt.Errorf("not found usecase %s. You need to create it first by call 'gogen usecase %s' ", usecaseName, usecaseName)
		}
	}

	for _, ot := range ds.Outports {

		file, err := os.Open(fmt.Sprintf("usecase/%s/port/outport.go", strings.ToLower(usecaseName)))
		if err != nil {
			return fmt.Errorf("not found usecase %s. You need to create it first by call 'gogen usecase %s' ", usecaseName, usecaseName)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		state := 0
		for scanner.Scan() {
			if state == 0 && strings.HasPrefix(scanner.Text(), fmt.Sprintf("type %sResponse struct {", ot.Name)) {
				state = 1
			} else //
			if state == 1 {
				if strings.HasPrefix(scanner.Text(), "}") {
					state = 2
					break
				} else {

					completeFieldWithType := strings.TrimSpace(scanner.Text())
					if len(completeFieldWithType) == 0 {
						continue
					}
					fieldWithType := strings.SplitN(completeFieldWithType, " ", 2)
					ot.ResponseFields = append(ot.ResponseFields, &NameType{
						Name: strings.TrimSpace(fieldWithType[0]),
					})

				}
			}
		}

		if state == 0 {
			return fmt.Errorf("not found usecase %s. You need to create it first by call 'gogen usecase %s' ", usecaseName, usecaseName)
		}
	}

	CreateFolder("datasource/%s", strings.ToLower(datasourceName))

	_ = WriteFileIfNotExist(
		"datasource/datasourceName/datasource._go",
		fmt.Sprintf("datasource/%s/%s.go", datasourceName, usecaseName),
		ds,
	)

	GoFormat(ds.PackagePath)

	return nil
}
