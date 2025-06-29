const _ = require("lodash");

var pkmnData = {};

const gbGames = ["red", "blue", "yellow", "gold", "silver", "crystal"]

const gbaGames = ["ruby", "sapphire", "emerald", "firered", "leafgreen"];

const dsGames = ["diamond", "pearl", "platinum", "heartgold", "soulsilver", "black", "white", "black-2", "white-2"];

const twoDGames = _.flattenDeep([gbGames, gbaGames, dsGames]);

function getLangEntries(arr, lang = "en") {
  return _.values(_.pickBy(arr, (item) => item.language.name === lang));
}

function getSingleLangEntry(arr, lang = "en") {
	return getLangEntries(arr, lang)[0];
}

function removeVersions(arr, omissions) {
  return _.values(_.omitBy(arr, (item) => omissions.includes(item.version.name)));
}

function stripDuplicateEntries(arr) {
  return _.uniqBy(arr, (item) => item.flavor_text.replace("\n", ""));
}

async function createAPIObj() {
  const dexLength = 1025;

  for (let i = 1; i <= dexLength; i++) {
    let p = await fetch(`https://pokeapi.co/api/v2/pokemon/${i}`)
    let s = await fetch(`https://pokeapi.co/api/v2/pokemon-species/${i}`)

    if (p && s) {
      let f = s.flavor_text_entries;
      f = getLangEntries(f, "en");
      f = removeVersions(f, twoDGames);
      f = stripDuplicateEntries(f)

      let g = s.genera;
      g = getSingleLangEntry(g, "en")

      let n = s.names;
      n = getSingleLangEntry(n, "ja-Hrkt")

      pkmnData.push({...p, flavor_text_entries: f, genera: g, name: n})
    }
  }
  console.log(pkmnData)
}

createAPIObj();