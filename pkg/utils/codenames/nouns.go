package codenames

// Mostly from https://github.com/lucasepe/codename

var nouns = []string{
	"abomination",
	"abyss",
	"agent",
	"amethyst",
	"amphibian",
	"andromeda",
	"annihilus",
	"anole",
	"anthem",
	"alchemist",
	"apocalypse",
	"aquagirl",
	"aquaman",
	"arachne",
	"arcade",
	"arcana",
	"archangel",
	"arclight",
	"ares",
	"argent",
	"arisia",
	"armadillo",
	"armor",
	"armory",
	"arrowette",
	"arsenal",
	"arsenic",
	"artemis",
	"artiee",
	"asgardian",
	"aspen",
	"atlas",
	"atom",
	"atomic",
	"avalanche",
	"azazel",
	"azrael",
	"aztec",
	"baguette",
	"crepe",
	"tartiflette",
	"fondue",
	"bourguignon",
	"rafale",
	"mirage",
	"napoleon",
	"cotedeboeuf",
	"concorde",
	"cassoulet",
	"escargots",
	"vin",
	"pinard",
	"absinthe",
	"calvados",
	"pastis",
	"cognac",
	"cotedurhone",
	"armagnac",
	"genepi",
	"cidre",
	"champagne",
	"raclette",
	"petanque",
	"grenouilles",
	"comte",
	"brillat",
	"cantal",
	"picodon",
	"pelardon",
	"chevre",
	"camembert",
	"brie",
	"reblochon",
	"munster",
	"roquefort",
	"bouillabaisse",
	"piedspaquets",
	"navette",
	"tapenade",
	"aioli",
	"calisson",
	"pompealhuile",
	"ratatouille",
	"pistou",
	"fougasse",
	"croissants",
	"surmulot",
	"painsuisse",
	"galette",
	"huitres",
	"ballistic",
	"banshee",
	"barb",
	"barbarella",
	"baroness",
	"barracuda",
	"bastion",
	"batgirl",
	"batman",
	"battle",
	"batwoman",
	"bazooka",
	"beak",
	"beast",
	"bebop",
	"becatron",
	"bedlam",
	"beef",
	"beetle",
	"bella",
	"belphegor",
	"bengal",
	"bette",
	"binary",
	"bionic",
	"bishop",
	"bizarro",
	"black",
	"blackbat",
	"blackheart",
	"blackout",
	"blade",
	"blastaar",
	"blindfold",
	"blink",
	"blitzkrieg",
	"blizzard",
	"blob",
	"blockbuster",
	"blok",
	"bloke",
	"blonde",
	"bloodaxe",
	"bloodberry",
	"bloodscream",
	"bloodstorm",
	"bloodstrike",
	"bloodhound",
	"bloom",
	"blossom",
	"blue",
	"bluestreak",
	"blur",
	"boom",
	"boomer",
	"boomerang",
	"booster",
	"bounty",
	"bubbles",
	"bug",
	"bulldozer",
	"bulleteer",
	"bulletgirl",
	"bullseye",
	"bumblebee",
	"burka",
	"burnout",
	"bushwacker",
	"buttercup",
	"butterfly",
	"cable",
	"calamity",
	"calendar",
	"caliban",
	"callisto",
	"calypso",
	"cammi",
	"cammy",
	"cannonball",
	"captain",
	"cardiac",
	"caretaker",
	"carnage",
	"cat",
	"catseye",
	"catwoman",
	"cerebro",
	"cha",
	"chameleon",
	"changeling",
	"chase",
	"chat",
	"cherry",
	"chimera",
	"choice",
	"chronomancer",
	"circuit",
	"cleopatra",
	"cloak",
	"clobber",
	"clobberella",
	"clover",
	"coagula",
	"cobra",
	"cobweb",
	"colleen",
	"colossus",
	"colt",
	"comedian",
	"comet",
	"conan",
	"constrictor",
	"contessa",
	"controller",
	"copperhead",
	"copycat",
	"cornelius",
	"corsair",
	"cosmo",
	"cottonmouth",
	"countess",
	"crane",
	"crazy",
	"crossbones",
	"crackmap",
	"crystal",
	"cyber",
	"cybergirl",
	"cyblade",
	"cyborg",
	"cyclone",
	"cyclops",
	"cypher",
	"cyslab",
	"dagger",
	"daredevil",
	"darkhawk",
	"darkstar",
	"darwin",
	"dawn",
	"dawnstar",
	"dazzler",
	"dead",
	"deadpool",
	"death",
	"deathbird",
	"deathcry",
	"deathlok",
	"deathstrike",
	"deep",
	"defenders",
	"demogoblin",
	"destine",
	"destiny",
	"devastator",
	"diablo",
	"diamond",
	"diamondback",
	"doctor",
	"doll",
	"dollar",
	"dolphin",
	"domino",
	"donatello",
	"doomsday",
	"doorman",
	"doppelganger",
	"dormammu",
	"dove",
	"dracula",
	"dragonfly",
	"dragonna",
	"drax",
	"dream",
	"dumb",
	"dusk",
	"dust",
	"dyna",
	"dynamite",
	"earthquake",
	"echo",
	"ego",
	"electra",
	"electro",
	"elektra",
	"elite",
	"elixir",
	"elongated",
	"empath",
	"empowered",
	"empress",
	"enchantress",
	"energizer",
	"epoch",
	"eradicator",
	"eternals",
	"eternity",
	"excalibur",
	"exodus",
	"expediter",
	"ezekiel",
	"fairchild",
	"faith",
	"falcon",
	"fallen",
	"famine",
	"fantomah",
	"fantomette",
	"fantomex",
	"fathom",
	"fenris",
	"feral",
	"fever",
	"fire",
	"firebird",
	"firebrand",
	"firedrake",
	"firefly",
	"firelord",
	"firestar",
	"firestorm",
	"fixer",
	"flaberella",
	"flamebird",
	"flash",
	"flatman",
	"flint",
	"flora",
	"forearm",
	"forerunner",
	"forge",
	"france",
	"freak",
	"free",
	"freefall",
	"frenzy",
	"fury",
	"galactus",
	"galvatron",
	"gambit",
	"gamora",
	"gangbuster",
	"ganymede",
	"garganta",
	"gargoyle",
	"gargoyles",
	"gateway",
	"gauntlet",
	"genesis",
	"ghost",
	"gladiator",
	"glitter",
	"glory",
	"goliath",
	"grandmaster",
	"graphics",
	"gravity",
	"greymalkin",
	"groot",
	"guardian",
	"guardsmen",
	"gunslinger",
	"gwen",
	"hairball",
	"hammerhead",
	"hardball",
	"harpoon",
	"haven",
	"havok",
	"hawk",
	"hawkeye",
	"hawkgirl",
	"hawkman",
	"hawkwoman",
	"heather",
	"hellboy",
	"hellcat",
	"hercules",
	"hiroim",
	"hitman",
	"hobgoblin",
	"holy",
	"hooded",
	"horridus",
	"howard",
	"hulk",
	"hulkling",
	"humbug",
	"huntara",
	"huntress",
	"husk",
	"hussar",
	"hydra",
	"hyperion",
	"ice",
	"iceman",
	"impulse",
	"indigo",
	"inertia",
	"infragirl",
	"inhumans",
	"ink",
	"insect",
	"invisible",
	"iron",
	"jackpot",
	"jigsaw",
	"joker",
	"jolt",
	"joystick",
	"jubilee",
	"judomaster",
	"juggernaut",
	"jungle",
	"juniper",
	"justice",
	"karate",
	"karatecha",
	"karma",
	"katana",
	"killmonger",
	"kinetix",
	"kingpin",
	"kitty",
	"klaw",
	"knockout",
	"komodo",
	"kree",
	"kronos",
	"lady",
	"ladyhawk",
	"lanolin",
	"laurel",
	"lavagirl",
	"layla",
	"leader",
	"leatherhead",
	"leatherneck",
	"legion",
	"leonardo",
	"leopardon",
	"lester",
	"lettuce",
	"liberty",
	"lifeguard",
	"ljpos",
	"lightning",
	"lightspeed",
	"lilandra",
	"lilith",
	"lime",
	"lionheart",
	"little",
	"lizard",
	"lockheed",
	"lockjaw",
	"longshot",
	"looker",
	"luckman",
	"maddog",
	"madripoor",
	"madrox",
	"maestro",
	"magik",
	"maginty",
	"magma",
	"magneto",
	"magus",
	"malice",
	"mandarin",
	"mandrill",
	"mandroid",
	"manhunter",
	"manitou",
	"manta",
	"mantis",
	"marionette",
	"marrow",
	"martian",
	"mastermind",
	"mathemanic",
	"mauler",
	"maximum",
	"maximus",
	"medusa",
	"megatron",
	"menace",
	"mentor",
	"mephisto",
	"metamorpho",
	"meteorite",
	"michaelangelo",
	"microbe",
	"microchip",
	"micromax",
	"midnight",
	"mighty",
	"mimic",
	"mindworm",
	"mint",
	"miracleman",
	"mirage",
	"misty",
	"mockingbird",
	"mongoose",
	"mongu",
	"monstress",
	"moondragon",
	"moonstar",
	"moonstone",
	"morbius",
	"mysterio",
	"mystique",
	"nebula",
	"negative",
	"nemesis",
	"neon",
	"netexec",
	"network",
	"nextwave",
	"night",
	"nightcat",
	"nightcrawler",
	"nighthawkc2",
	"nightmare",
	"nightshade",
	"nightstar",
	"nightveil",
	"nightwing",
	"nitro",
	"nmap",
	"nocturne",
	"nomad",
	"northstar",
	"nova",
	"nuke",
	"odin",
	"ogun",
	"onslaught",
	"onyx",
	"oracle",
	"orion",
	"overlord",
	"owl",
	"owlman",
	"owlwoman",
	"paladin",
	"pandemic",
	"pantha",
	"parasite",
	"patch",
	"patriot",
	"payback",
	"penance",
	"penguin",
	"pestilence",
	"phalanx",
	"phantom",
	"phoenix",
	"photon",
	"piledriver",
	"pigeon",
	"plastic",
	"poison",
	"polaris",
	"post",
	"power",
	"princess",
	"prism",
	"prodigy",
	"psylocke",
	"punisher",
	"purifiers",
	"pyro",
	"quasar",
	"queen",
	"quicksilver",
	"rage",
	"raider",
	"rainbow",
	"rainmaker",
	"rampage",
	"random",
	"rafff",
	"raisin",
	"raptor",
	"rapture",
	"reaper",
	"redwing",
	"reptil",
	"rescue",
	"revanche",
	"reverse",
	"rhino",
	"ricochet",
	"rictor",
	"riddler",
	"risque",
	"rocket",
	"rockslide",
	"rogue",
	"sailor",
	"sandman",
	"saracen",
	"sasquatch",
	"satana",
	"saturn",
	"sauron",
	"savant",
	"scalphunter",
	"scarecrow",
	"scarlet",
	"scorpion",
	"scourge",
	"scrambler",
	"scream",
	"screwball",
	"secret",
	"sentinel",
	"sentinels",
	"sentry",
	"sepulchre",
	"serpentor",
	"shadow",
	"shadowcat",
	"shadoweyes",
	"shaman",
	"shamrock",
	"shocker",
	"shockwave",
	"shotgun",
	"shredder",
	"shriek",
	"shrinking",
	"siege",
	"silhouette",
	"silver",
	"silverclaw",
	"silvermane",
	"siren",
	"skullbuster",
	"skyrocket",
	"slapstick",
	"slayback",
	"sleeper",
	"sleepwalker",
	"slipstream",
	"smasher",
	"snowbird",
	"songbird",
	"spartan",
	"spectrum",
	"speedball",
	"speedy",
	"spellbinder",
	"sphinx",
	"spider",
	"spiral",
	"spirit",
	"spitfire",
	"spoiler",
	"spot",
	"sprite",
	"spy",
	"spyke",
	"squirrel",
	"starbolt",
	"stardust",
	"starfire",
	"starfox",
	"stargirl",
	"starhawk",
	"starwoman",
	"sliver",
	"steel",
	"stinger",
	"stingray",
	"storm",
	"stormtrooper",
	"stranger",
	"stripperella",
	"stryfe",
	"stunner",
	"sunfire",
	"sunspot",
	"supergirl",
	"supergran",
	"superman",
	"supernaut",
	"superwoman",
	"swift",
	"switch",
	"swordsman",
	"synch",
	"tag",
	"talisman",
	"talkback",
	"talon",
	"talos",
	"tank",
	"tara",
	"tarantula",
	"tarot",
	"taskmaster",
	"tattoo",
	"tecna",
	"tempest",
	"tenebrous",
	"terror",
	"terry",
	"thing",
	"thunder",
	"thunderball",
	"thunderbird",
	"thunderbolt",
	"tiger",
	"timeslip",
	"tinkerer",
	"titaness",
	"titania",
	"toad",
	"tombstone",
	"toxin",
	"trauma",
	"triathlon",
	"triceraton",
	"triplicate",
	"triton",
	"tsunami",
	"turbo",
	"tyrannus",
	"ultimatum",
	"ultimo",
	"ultra",
	"ultragirl",
	"ultrawoman",
	"ultron",
	"unicorn",
	"valkyrie",
	"vampirella",
	"vampiro",
	"vanisher",
	"vapor",
	"vector",
	"velocity",
	"vengeance",
	"venom",
	"venus",
	"vermin",
	"vigilante",
	"vindicator",
	"violations",
	"violator",
	"violet",
	"viper",
	"virtuous",
	"vision",
	"vivisector",
	"vixen",
	"vogue",
	"void",
	"voodoo",
	"vulcan",
	"vulture",
	"wawawa",
	"wallflower",
	"wallen",
	"wallow",
	"warbird",
	"warbound",
	"warhawk",
	"warlock",
	"warpath",
	"warstar",
	"wasp",
	"watchmen",
	"wendigo",
	"whirlwind",
	"whistler",
	"whizzer",
	"wiccan",
	"widget",
	"wild",
	"wildcat",
	"winged",
	"witchblade",
	"witchfire",
	"wolfpack",
	"wolfsbane",
	"wolverine",
	"wonder",
	"wraith",
	"wrecker",
	"yellowjacket",
	"zombie",
}
