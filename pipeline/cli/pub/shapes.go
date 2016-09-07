package pub

import (
	"github.com/naveego/api/pipeline/publisher"
	"github.com/spf13/cobra"
)

var shapesCmd = &cobra.Command{
	Use:   "shapes",
	Short: "Gets the shapes that a publisher can publish to the pipeline",
	RunE: func(cmd *cobra.Command, args []string) error {
		pubFactory, err := publisher.GetFactory(TypeName)
		if err != nil {
			return err
		}

		ctx := publisher.Context{}
		p := pubFactory()
		_, err = p.Shapes(ctx)
		return err
	},
}
