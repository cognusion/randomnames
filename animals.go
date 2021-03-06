package randomnames

// List of animals from https://gist.githubusercontent.com/atduskgreg/3cf8ef48cb0d29cf151bedad81553a54/raw/82f142562cf50b0f6fb8010f890b2f934093553e/animals.txt

import (
	"math/rand"
	"sync"
)

func init() {
	animalSize = len(Animals)
}

// RandomAnimal returns a pseudo-random animal from the list
func RandomAnimal() string {
	return Animals[rand.Intn(animalSize)]
}

// SafeRandomAnimal returns a pseudo-random animal from the list goro-safely
func SafeRandomAnimal() string {
	nameSafety.Lock()
	s := RandomAnimal()
	nameSafety.Unlock() // manual unlock so we don't incur the defer penalty
	return s
}

var (
	animalSafety sync.Mutex
	animalSize   int
	Animals      []string = []string{
		"Aardvark",
		"Aardwolf",
		"Albatross",
		"Alligator",
		"Alpaca",
		"Alpaca",
		"Amphibian",
		"Anaconda",
		"Angelfish",
		"Anglerfish",
		"Ant",
		"Anteater",
		"Antelope",
		"Antlion",
		"Ape",
		"Aphid",
		"Arctic Fox",
		"Arctic Wolf",
		"Armadillo",
		"Arrow crab",
		"Asp",
		"Ass",
		"Baboon",
		"Bactrian camel",
		"Badger",
		"Bald eagle",
		"Bali cattle",
		"Bandicoot",
		"Barnacle",
		"Barracuda",
		"Basilisk",
		"Bass",
		"Bat",
		"Beaked whale",
		"Bear",
		"Beaver",
		"Bedbug",
		"Bee",
		"Beetle",
		"Bird",
		"Bison",
		"Blackbird",
		"Black panther",
		"Black widow spider",
		"Blue bird",
		"Blue jay",
		"Blue whale",
		"Boa",
		"Boar",
		"Bobcat",
		"Bobolink",
		"Bonobo",
		"Booby",
		"Bovid",
		"Box jellyfish",
		"Buffalo",
		"Bug",
		"Butterfly",
		"Buzzard",
		"Camel",
		"Canid",
		"Canidae",
		"Cape buffalo",
		"Capybara",
		"Cardinal",
		"Caribou",
		"Carp",
		"Cat",
		"Caterpillar",
		"Catfish",
		"Catshark",
		"Cattle",
		"Centipede",
		"Cephalopod",
		"Chameleon",
		"Cheetah",
		"Chickadee",
		"Chicken",
		"Chimpanzee",
		"Chinchilla",
		"Chipmunk",
		"Cicada",
		"Clam",
		"Clownfish",
		"Cobra",
		"Cockroach",
		"Cod",
		"Condor",
		"Constrictor",
		"Coral",
		"Cougar",
		"Cow",
		"Coyote",
		"Crab",
		"Crane",
		"Crane fly",
		"Crawdad",
		"Crayfish",
		"Cricket",
		"Crocodile",
		"Crow",
		"Cuckoo",
		"Damselfly",
		"Deer",
		"Dingo",
		"Dinosaur",
		"Dog",
		"Dolphin",
		"Donkey",
		"Donkey",
		"Donkey",
		"Dormouse",
		"Dove",
		"Dragon",
		"Dragonfly",
		"Dromedary camel",
		"Duck",
		"Dung beetle",
		"Eagle",
		"Earthworm",
		"Earwig",
		"Echidna",
		"Eel",
		"Egret",
		"Elephant",
		"Elephant seal",
		"Elk",
		"Emu",
		"English pointer",
		"Ermine",
		"Falcon",
		"Fancy mouse",
		"Fancy rat",
		"Felidae",
		"Ferret",
		"Finch",
		"Firefly",
		"Fish",
		"Flamingo",
		"Flea",
		"Fly",
		"Flyingfish",
		"Fowl",
		"Fox",
		"Frog",
		"Fruit bat",
		"Galliform",
		"Gamefowl",
		"Gayal",
		"Gazelle",
		"Gecko",
		"Gerbil",
		"Giant panda",
		"Giant squid",
		"Gibbon",
		"Gila monster",
		"Giraffe",
		"Goat",
		"Goldfish",
		"Goose",
		"Gopher",
		"Gorilla",
		"Grasshopper",
		"Great blue heron",
		"Great white shark",
		"Grizzly bear",
		"Ground shark",
		"Ground sloth",
		"Grouse",
		"Guan",
		"Guanaco",
		"Guineafowl",
		"Guinea pig",
		"Gull",
		"Guppy",
		"Haddock",
		"Halibut",
		"Hammerhead shark",
		"Hamster",
		"Hare",
		"Harrier",
		"Hawk",
		"Hedgehog",
		"Hermit crab",
		"Heron",
		"Herring",
		"Hippopotamus",
		"Hookworm",
		"Hornet",
		"Horse",
		"Hoverfly",
		"Hummingbird",
		"Humpback whale",
		"Hyena",
		"Iguana",
		"Impala",
		"Irukandji jellyfish",
		"Jackal",
		"Jaguar",
		"Jay",
		"Jellyfish",
		"Junglefowl",
		"Kangaroo",
		"Kangaroo mouse",
		"Kangaroo rat",
		"Kingfisher",
		"Kite",
		"Kiwi",
		"Koala",
		"Koi",
		"Komodo dragon",
		"Krill",
		"Lab rat",
		"Ladybug",
		"Lamprey",
		"Landfowl",
		"Land snail",
		"Lark",
		"Leech",
		"Lemming",
		"Lemur",
		"Leopard",
		"Leopon",
		"Limpet",
		"Lion",
		"Lizard",
		"Llama",
		"Lobster",
		"Locust",
		"Loon",
		"Louse",
		"Lungfish",
		"Lynx",
		"Macaw",
		"Mackerel",
		"Magpie",
		"Mammal",
		"Manatee",
		"Mandrill",
		"Manta ray",
		"Marlin",
		"Marmoset",
		"Marmot",
		"Marsupial",
		"Marten",
		"Mastodon",
		"Meadowlark",
		"Meerkat",
		"Mink",
		"Minnow",
		"Mite",
		"Mockingbird",
		"Mole",
		"Mollusk",
		"Mongoose",
		"Monitor lizard",
		"Monkey",
		"Moose",
		"Mosquito",
		"Moth",
		"Mountain goat",
		"Mouse",
		"Mule",
		"Muskox",
		"Narwhal",
		"Newt",
		"New World quail",
		"Nightingale",
		"Ocelot",
		"Octopus",
		"Old World quail",
		"Opossum",
		"Orangutan",
		"Orca",
		"Ostrich",
		"Otter",
		"Owl",
		"Ox",
		"Panda",
		"Panther",
		"Panthera hybrid",
		"Parakeet",
		"Parrot",
		"Parrotfish",
		"Partridge",
		"Peacock",
		"Peafowl",
		"Pelican",
		"Penguin",
		"Perch",
		"Peregrine falcon",
		"Pheasant",
		"Pig",
		"Pigeon",
		"Pike",
		"Pilot whale",
		"Pinniped",
		"Piranha",
		"Planarian",
		"Platypus",
		"Polar bear",
		"Pony",
		"Porcupine",
		"Porpoise",
		"Portuguese man o' war",
		"Possum",
		"Prairie dog",
		"Prawn",
		"Praying mantis",
		"Primate",
		"Ptarmigan",
		"Puffin",
		"Puma",
		"Python",
		"Quail",
		"Quelea",
		"Quokka",
		"Rabbit",
		"Raccoon",
		"Rainbow trout",
		"Rat",
		"Rattlesnake",
		"Raven",
		"Ray",
		"Red panda",
		"Reindeer",
		"Reptile",
		"Rhinoceros",
		"Right whale",
		"Ringneck dove",
		"Roadrunner",
		"Rodent",
		"Rook",
		"Rooster",
		"Roundworm",
		"Saber-toothed cat",
		"Sailfish",
		"Salamander",
		"Salmon",
		"Sawfish",
		"Scale insect",
		"Scallop",
		"Scorpion",
		"Seahorse",
		"Sea lion",
		"Sea slug",
		"Sea snail",
		"Shark",
		"Sheep",
		"Shrew",
		"Shrimp",
		"Siamese fighting fish",
		"Silkworm",
		"Silverfish",
		"Skink",
		"Skunk",
		"Sloth",
		"Slug",
		"Smelt",
		"Snail",
		"Snake",
		"Snipe",
		"Snow leopard",
		"Society finch",
		"Sockeye salmon",
		"Sole",
		"Sparrow",
		"Sperm whale",
		"Spider",
		"Spider monkey",
		"Spoonbill",
		"Squid",
		"Squirrel",
		"Starfish",
		"Star-nosed mole",
		"Steelhead trout",
		"Stingray",
		"Stoat",
		"Stork",
		"Sturgeon",
		"Sugar glider",
		"Swallow",
		"Swan",
		"Swift",
		"Swordfish",
		"Swordtail",
		"Tahr",
		"Takin",
		"Tapir",
		"Tarantula",
		"Tarsier",
		"Tasmanian devil",
		"Termite",
		"Tern",
		"Thrush",
		"Tick",
		"Tiger",
		"Tiger shark",
		"Tiglon",
		"Toad",
		"Tortoise",
		"Toucan",
		"Trapdoor spider",
		"Tree frog",
		"Trout",
		"Tuna",
		"Turkey",
		"Turtle",
		"Tyrannosaurus",
		"Urial",
		"Vampire bat",
		"Vampire squid",
		"Vicuna",
		"Viper",
		"Vole",
		"Vulture",
		"Wallaby",
		"Walrus",
		"Warbler",
		"Wasp",
		"Water Boa",
		"Water buffalo",
		"Weasel",
		"Whale",
		"Whippet",
		"Whitefish",
		"Whooping crane",
		"Wildcat",
		"Wildebeest",
		"Wildfowl",
		"Wolf",
		"Wolverine",
		"Wombat",
		"Woodpecker",
		"Worm",
		"Wren",
		"Xerinae",
		"X-ray fish",
		"Yak",
		"Yellow perch",
		"Zebra",
		"Zebra finch",
	}
)
