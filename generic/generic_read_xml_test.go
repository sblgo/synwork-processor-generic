package generic

import (
	"os"
	"strings"
	"testing"

	"sbl.systems/go/synwork/plugin-sdk/tunit"
)

func TestGenericReadXml01(t *testing.T) {
	_defs := `
	method "read-xml" "dum" "xml_struct" {
		file_name = "<FILENAME>"
	}
	`
	xml := `<xml>
		<row><a>A1</a><b>B1</b></row>
		<row><a>A2</a><b>B2</b></row>
		<row><a>A3</a><b>B3</b></row>
		<row><a>A4</a><b>B4</b></row>
	</xml>`

	file, err := os.CreateTemp(".", "generic_read_xml")
	if err != nil {
		t.Fatalf("can't create tmp file %s", err.Error())
	}
	defer os.Remove(file.Name())
	_, err = file.WriteString(xml)
	if err != nil {
		t.Fatalf("can't write tmp file %s", err.Error())
	}
	_defs = strings.ReplaceAll(_defs, "<FILENAME>", file.Name())
	mm := tunit.MethodMock{
		ProcessorDef: Opts.Provider,
		InstanceMock: struct{}{},
		ExecFunc:     generic_read_xml,
		References:   map[string]interface{}{},
	}
	result := tunit.CallMockMethod(t, mm, _defs)
	xmlRoot := result["xml"].(map[string]interface{})
	rows := xmlRoot["row"].([]interface{})
	if len(rows) != 4 {
		t.Fatal()
	}
}

func TestGenericReadXml02(t *testing.T) {
	_defs := `
	method "read-xml" "dum" "xml_struct" {
		file_name = "<FILENAME>"
		arrays = "/xml/row"
	}
	`
	xml := `<xml>
		<row><a>A1</a><b>B1</b></row>
	</xml>`

	file, err := os.CreateTemp(".", "generic_read_xml")
	if err != nil {
		t.Fatalf("can't create tmp file %s", err.Error())
	}
	defer os.Remove(file.Name())
	_, err = file.WriteString(xml)
	if err != nil {
		t.Fatalf("can't write tmp file %s", err.Error())
	}
	_defs = strings.ReplaceAll(_defs, "<FILENAME>", file.Name())
	mm := tunit.MethodMock{
		ProcessorDef: Opts.Provider,
		InstanceMock: struct{}{},
		ExecFunc:     generic_read_xml,
		References:   map[string]interface{}{},
	}
	result := tunit.CallMockMethod(t, mm, _defs)
	xmlRoot := result["xml"].(map[string]interface{})
	rows := xmlRoot["row"].([]interface{})
	if len(rows) != 1 {
		t.Fatal()
	}
}

func TestGenericReadXml03(t *testing.T) {
	_defs := `
	method "read-xml" "dum" "xml_struct" {
		file_name = "<FILENAME>"
		arrays = "/xml/row"
	}
	`
	xml := test03

	file, err := os.CreateTemp(".", "generic_read_xml")
	if err != nil {
		t.Fatalf("can't create tmp file %s", err.Error())
	}
	defer os.Remove(file.Name())
	_, err = file.WriteString(xml)
	if err != nil {
		t.Fatalf("can't write tmp file %s", err.Error())
	}
	_defs = strings.ReplaceAll(_defs, "<FILENAME>", file.Name())
	mm := tunit.MethodMock{
		ProcessorDef: Opts.Provider,
		InstanceMock: struct{}{},
		ExecFunc:     generic_read_xml,
		References:   map[string]interface{}{},
	}
	result := tunit.CallMockMethod(t, mm, _defs)
	xmlRoot := result["xml"].(map[string]interface{})
	rows := xmlRoot["row"].([]interface{})
	if len(rows) != 1 {
		t.Fatal()
	}
}

const test03 = `<?xml version="1.0" encoding="utf-8"?>
<ns:MT_OUT_Confirmation_PPL3xx_MES xmlns:ns="http:AGILE_PP_WEBER_SE_Z_PRODORDCONF_GI">
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269594</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200811007</MATNR>
		<MENGE>456.5</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.61</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269595</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200811008</MATNR>
		<MENGE>571.0</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.613</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269596</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200811009</MATNR>
		<MENGE>541.5</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.62</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269597</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200811064</MATNR>
		<MENGE>587.0</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.623</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269598</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200863467</MATNR>
		<MENGE>0.0</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.63</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269599</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200811310</MATNR>
		<MENGE>334.0</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.633</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269600</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200860777</MATNR>
		<MENGE>84.5</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.64</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269601</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200845292</MATNR>
		<MENGE>196.5</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.643</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269602</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200860775</MATNR>
		<MENGE>23.5</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.65</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269603</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200811476</MATNR>
		<MENGE>43.0</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.653</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269604</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200859697</MATNR>
		<MENGE>2.84</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.66</TIMESTAMPv>
	</row>
	<row>
		<ID></ID>
		<SATZA></SATZA>
		<WERKS></WERKS>
		<AUFNR></AUFNR>
		<VORNR></VORNR>
		<ARBID></ARBID>
		<LMNGA></LMNGA>
		<XMNGA></XMNGA>
		<MEINS></MEINS>
		<GRUND></GRUND>
		<ERDAT></ERDAT>
		<ERTIM></ERTIM>
		<BUDAT></BUDAT>
		<STATUS></STATUS>
		<TIMESTAMP></TIMESTAMP>
		<IDv>6269605</IDv>
		<AUFNRv>63096264</AUFNRv>
		<BATCHNO>110004</BATCHNO>
		<WERKSv>L333</WERKSv>
		<MATNR>5200811486</MATNR>
		<MENGE>1.4</MENGE>
		<MEINSv>KGM</MEINSv>
		<LGORT>1001</LGORT>
		<BWART>261</BWART>
		<STATUSv>1</STATUSv>
		<TIMESTAMPv>2022-02-08 14:24:05.663</TIMESTAMPv>
	</row>
</ns:MT_OUT_Confirmation_PPL3xx_MES>`
