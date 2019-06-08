package cmd

import (
	"fmt"

	"github.com/identixone/identixone-go/api/client"
	"github.com/identixone/identixone-go/api/const/conf"
	"github.com/identixone/identixone-go/api/person"
	"github.com/identixone/identixone-go/utils"
	"github.com/spf13/cobra"
)

const (
	personaCreate      = "create"
	personaSearch      = "search"
	personaDelete      = "delete"
	personaReinitImage = "reinit-image"
	personaReinitId    = "reinit-id"
)

var (
	photoPath          string
	liveness           bool
	livenessOnly       bool
	asm                bool
	reinitId           int
	createOnHa         bool
	createOnJunk       bool
	createOnlyLiveness bool
	reinitConf         string
)

var personsCmd = &cobra.Command{
	Use: "persons [command]",
	//Aliases: []string{"command"},
	Short: "working with personas",
	Long: `Personas are primary patterns of faces uploaded to the platform either via user upload or automatically from cameras. 
Effectively, this is the same tag as sources, but whereas sources show where photos come from, personas tag who is in these photos.

Available Commands:
	create			creating personas from photos
	search			searching personas from photos
	delete			searching personas from photos
	reinit-image		re-initialization from a photo
	reinit-id		re-initialization from a entry`,
	Example: `Create person if liveness is passed:
	identixone persons create --photo img/v2885.png --source default --liveness-only true

Searching person with liveness and asm:
	identixone persons search --photo img/v2885.png --liveness true --asm true
`,
	ValidArgs: []string{personaCreate, personaSearch, personaDelete, personaReinitImage, personaReinitId},
	Args:      cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case personaCreate:
			if photoPath == "" {
				printAndExit("photo is required.")
			}
			if sourceName == "" {
				printAndExit("source is required.")
			}
			cp, err := person.NewPersonaCreateRequest(photoPath, sourceName)
			ifErrorExit(err)

			cp.SetAsm(asm).
				SetLiveness(liveness).
				SetFacesize(faceSize).
				SetCreateLivenessOnly(createOnlyLiveness).
				SetCreateOnHa(createOnHa).
				SetCreateOnJunk(createOnJunk)
			//fmt.Printf("%+v\n", cp)
			c, err := client.NewClient()
			ifErrorExit(err)
			person, err := c.Persons().Create(cp)
			ifErrorExit(err)
			ifErrorExit(utils.PrettyPrint(person))
			break
		case personaSearch:
			if photoPath == "" {
				printAndExit("photo is required.")
			}

			sch, err := person.NewSearch(photoPath, asm, liveness)
			ifErrorExit(err)

			c, err := client.NewClient()
			ifErrorExit(err)
			result, err := c.Persons().Search(sch)
			ifErrorExit(err)
			ifErrorExit(utils.PrettyPrint(result))
			break
		case personaDelete:
			if idxid == "" {
				printAndExit("missing idxid flag.")
			}
			c, err := client.NewClient()
			ifErrorExit(err)
			ifErrorExit(c.Persons().Delete(idxid))
			fmt.Println("person successfully deleted.")
		case personaReinitId:
			if reinitId == 0 {
				printAndExit("missing id for reinit.")
			}
			req := person.ReinitIdRequest{Id: reinitId}
			if faceSize != 0 {
				req.Facesize = faceSize
			}
			c, err := client.NewClient()
			ifErrorExit(err)
			ifErrorExit(c.Persons().ReinitId(req))
			fmt.Println("reinit success.")
		case personaReinitImage:
			if photoPath == "" {
				printAndExit("photo is required.")
			}

			if idxid == "" {
				printAndExit("idxid is required.")
			}

			req, err := person.NewReinitImageRequest(photoPath, idxid)
			ifErrorExit(err)

			req.SetFacesize(faceSize).SetLiveness(liveness).SetReinitLivenessOnly(livenessOnly)

			if reinitConf != "" {
				req.SetConf(conf.Conf(reinitConf))
			}
			c, err := client.NewClient()
			ifErrorExit(err)
			ifErrorExit(c.Persons().ReinitImage(req))
			fmt.Println("reinit success.")

		}
	},
}

func init() {
	personsCmd.Flags().StringVarP(&photoPath, "photo", "p", "", "path to photo (png or jpeg file format) for searching or creating")
	personsCmd.Flags().StringVarP(&sourceName, "source", "s", "", "source name for persona creating")
	personsCmd.Flags().BoolVar(&liveness, "liveness", false, "defines settings for image liveness check (default false)")
	personsCmd.Flags().BoolVar(&livenessOnly, "liveness-only", false, "reinit only if photo is liveness (default false)")
	personsCmd.Flags().BoolVar(&asm, "asm", false, "defines settings for age/sex/mood check (default false)")
	personsCmd.Flags().BoolVar(&createOnHa, "create-on-ha", false, "a command to create a persona, even though there is already one persona with ha (high accuracy) result in the database (default false)")
	personsCmd.Flags().BoolVar(&createOnJunk, "create-on-junk", false, "a command to create a persona, even though there is already one persona with junk result in the database (default false)")
	personsCmd.Flags().StringVar(&idxid, "idxid", "", "idxid for manipulating with persona")
	personsCmd.Flags().IntVar(&reinitId, "id", 0, "id for reinit-id")
	personsCmd.Flags().IntVar(&faceSize, "facesize", 0, "minimum face square size in pixels")
	personsCmd.Flags().StringVar(&reinitConf, "conf", "ha", "minimum result of comparison between primary photo and uploaded photo")

	rootCmd.AddCommand(personsCmd)
}
