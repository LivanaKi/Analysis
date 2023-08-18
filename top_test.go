package analysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `It goes without saying that humans (mammals identifiable as those that stand upright and are comparatively advanced and capable of detailed thought) have pretty remarkable bodies, given all that they've accomplished. (Furthermore, an especially intelligent human brain produced this text!) To be sure, humans have overcome predators, disease, and all sorts of other obstacles over thousands of years.

To fully understand and appreciate these accomplishments, let's take at some of the most well-known parts of the human body!

The head, or the spherical body part that contains the brain and rests at the top of the human body, has quite a few individual organs and body parts on it. (It should quickly be mentioned that hair occupies the space on top of the head, and the ears, the organs responsible for hearing, are located on either side of the head.) From top to bottom, the eyebrows, or horizontal strips of hair that can be found above the eye, are the first components of the head. The eyes are below them, and are round, orb-like organs that allow humans to see.

The eyes make way for the nose, or an external (sticking-out) organ that plays an important part in the breathing and bacteria-elimination processes. Below that is the mouth, or a wide, cavernous organ that chews food, removes bacteria, helps with breathing, and more. The mouth contains teeth, or small, white-colored, pointed body parts used to chew food, and the tongue, or a red-colored, boneless organ used to chew food and speak.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {

				expected := []string{
					"the",
					"and",
					"of",
					"that",
					"or",
					"to",
					"are",
					"body",
					"head",
					"a",
				}
			
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				    "that",    
				    "the",
					"and",
					"of",
					"to",
					"or",
					"are",
					"body",
					"head",
					"a",
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}
