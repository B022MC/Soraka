import { B as ByID } from "./index-DR7AVCif.js";
function Any(source) {
  return (
    /** @type {T} */
    source
  );
}
function Array(element) {
  if (element === Any) {
    return (source) => source === null ? [] : source;
  }
  return (source) => {
    if (source === null) {
      return [];
    }
    for (let i = 0; i < source.length; i++) {
      source[i] = element(source[i]);
    }
    return source;
  };
}
function Nullable(element) {
  if (element === Any) {
    return Any;
  }
  return (source) => source === null ? null : element(source);
}
class MatchBrief {
  /** Creates a new MatchBrief instance. */
  constructor($$source = {}) {
    if (!("id" in $$source)) {
      this["id"] = 0;
    }
    if (!("result" in $$source)) {
      this["result"] = "";
    }
    if (!("mode" in $$source)) {
      this["mode"] = "";
    }
    if (!("mode_detail" in $$source)) {
      this["mode_detail"] = "";
    }
    if (!("kills" in $$source)) {
      this["kills"] = 0;
    }
    if (!("deaths" in $$source)) {
      this["deaths"] = 0;
    }
    if (!("assists" in $$source)) {
      this["assists"] = 0;
    }
    if (!("cs" in $$source)) {
      this["cs"] = 0;
    }
    if (!("gold" in $$source)) {
      this["gold"] = 0;
    }
    if (!("time" in $$source)) {
      this["time"] = "";
    }
    if (!("level" in $$source)) {
      this["level"] = 0;
    }
    if (!("champion" in $$source)) {
      this["champion"] = "";
    }
    if (!("spells" in $$source)) {
      this["spells"] = [];
    }
    if (!("items" in $$source)) {
      this["items"] = [];
    }
    if (!("map" in $$source)) {
      this["map"] = "";
    }
    Object.assign(this, $$source);
  }
  /**
   * Creates a new MatchBrief instance from a string or object.
   */
  static createFrom($$source = {}) {
    const $$createField12_0 = $$createType0$1;
    const $$createField13_0 = $$createType0$1;
    let $$parsedSource = typeof $$source === "string" ? JSON.parse($$source) : $$source;
    if ("spells" in $$parsedSource) {
      $$parsedSource["spells"] = $$createField12_0($$parsedSource["spells"]);
    }
    if ("items" in $$parsedSource) {
      $$parsedSource["items"] = $$createField13_0($$parsedSource["items"]);
    }
    return new MatchBrief($$parsedSource);
  }
}
class RankSummary {
  /** Creates a new RankSummary instance. */
  constructor($$source = {}) {
    if (!("type" in $$source)) {
      this["type"] = "";
    }
    if (!("total" in $$source)) {
      this["total"] = 0;
    }
    if (!("rate" in $$source)) {
      this["rate"] = "";
    }
    if (!("wins" in $$source)) {
      this["wins"] = 0;
    }
    if (!("loses" in $$source)) {
      this["loses"] = 0;
    }
    if (!("rank" in $$source)) {
      this["rank"] = "";
    }
    if (!("lp" in $$source)) {
      this["lp"] = 0;
    }
    if (!("seasonMax" in $$source)) {
      this["seasonMax"] = "";
    }
    if (!("lastSeason" in $$source)) {
      this["lastSeason"] = "";
    }
    Object.assign(this, $$source);
  }
  /**
   * Creates a new RankSummary instance from a string or object.
   */
  static createFrom($$source = {}) {
    let $$parsedSource = typeof $$source === "string" ? JSON.parse($$source) : $$source;
    return new RankSummary($$parsedSource);
  }
}
const $$createType0$1 = Array(Any);
function GetClientPath() {
  let $resultPromise = ByID(4150102446);
  return $resultPromise;
}
function GetRankSummary() {
  let $resultPromise = ByID(4148630304);
  let $typingPromise = $resultPromise.then(($result) => {
    return $$createType2($result);
  });
  $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
  return $typingPromise;
}
function ListRecentMatches(limit) {
  let $resultPromise = ByID(1081314318, limit);
  let $typingPromise = $resultPromise.then(($result) => {
    return $$createType5($result);
  });
  $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
  return $typingPromise;
}
function StartClient() {
  let $resultPromise = ByID(445256185);
  return $resultPromise;
}
const $$createType0 = RankSummary.createFrom;
const $$createType1 = Nullable($$createType0);
const $$createType2 = Array($$createType1);
const $$createType3 = MatchBrief.createFrom;
const $$createType4 = Nullable($$createType3);
const $$createType5 = Array($$createType4);
export {
  GetClientPath as G,
  ListRecentMatches as L,
  StartClient as S,
  GetRankSummary as a
};
