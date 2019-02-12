package configurations

import "os"

var Port, LineFileFullpath, SDLEndpoint string

func init() {
	if LineFileFullpath = os.Getenv("LINE_INGESTOR_LINE_FILE_URI"); LineFileFullpath == "" {
		LineFileFullpath = "./file/SRV_LINEAS.DAT"
	}

	if Port = os.Getenv("LINE_INGESTOR_SERVER_PORT"); Port == "" {
		Port = "8080"
	}

	if SDLEndpoint = os.Getenv("LINE_INGESTOR_SDL_ENDPOINT"); SDLEndpoint == "" {
		panic("SDL endpoint has not been set.")
	}
}
