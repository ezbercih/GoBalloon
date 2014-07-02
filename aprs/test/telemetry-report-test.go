package main

import (
	"fmt"
	"github.com/chrissnell/GoBalloon/aprs"
	"log"
)

func main() {
	s := aprs.StdTelemetryReport{
		Sequence: 1,
		A1:       206,
		A2:       100,
		A3:       20,
		A4:       36,
		A5:       40,
		Digital:  77,
	}

	c := aprs.CompressedTelemetryReport{
		Sequence: 3,
		A1:       7740,
		A2:       105,
		A3:       20,
		A4:       3006,
		A5:       403,
		Digital:  77,
	}

	tr := aprs.CreateUncompressedTelemetryReport(&s)
	fmt.Printf("Standard telemetry structure: %v\n", tr)
	p, err := aprs.ParseUncompressedTelemetryReport(tr)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Printf("Standard telemetry report%v\n", p)

	fmt.Printf("Compressed telemetry structure: %+v\n", c)
	ctr, err := aprs.CreateCompressedTelemetryReport(&c)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fmt.Printf("Compressed telemetry report: %v\n", ctr)

	pc, err := aprs.ParseCompressedTelemetryReport(ctr)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Printf("Decompressed compressed telemetry report: %+v\n", pc)

}
