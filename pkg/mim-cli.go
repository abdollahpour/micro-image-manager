package main

/*func parseProfile(profile string) (*processor.Profile, error) {
	p1 := strings.Split(profile, "_")
	p2 := strings.Split(p1[1], "x")

	width, err := strconv.Atoi(p2[0])
	if err != nil {
		return nil, err
	}
	height, err := strconv.Atoi(p2[1])
	if err != nil {
		return nil, err
	}

	return &processor.Profile{
		Name:   p1[0],
		Width:  width,
		Height: height,
	}, nil
}

func main() {
	imagePtr := flag.String("image", "", "image filename")
	distPtr := flag.String("dist", "", "image output directory")
	flag.Parse()

	if *imagePtr == "" || *distPtr == "" || len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide required parameters")
		fmt.Fprintln(os.Stderr, "Example: mim -image=input.jpg -dist=output large_800x600")
		flag.PrintDefaults()
		os.Exit(1)
	}

	id := uuid.NewString()

	localStorage := storage.NewLocalStorage(*distPtr)
	imageProcessor := processor.NewBimgProcessor("temp")

	for _, profileArg := range flag.Args() {
		profile, err := parseProfile(profileArg)

		imageProcessor.Process(id, )

		localStorage.Store("", p1, "")
	}

	fmt.Println(*imagePtr)
}


func receive(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	event.Data()
}

func main() {
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), receive))
}
*/
