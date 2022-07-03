package lib

type OxfordResponse struct {
  ID       string `json:"id"`
  Metadata struct {
    Operation string `json:"operation"`
    Provider  string `json:"provider"`
    Schema    string `json:"schema"`
  } `json:"metadata"`
  Results []struct {
    ID             string `json:"id"`
    Language       string `json:"language"`
    LexicalEntries []struct {
      Entries []struct {
        Etymologies []string `json:"etymologies"`
        Inflections []struct {
          GrammaticalFeatures []struct {
            ID   string `json:"id"`
            Text string `json:"text"`
            Type string `json:"type"`
          } `json:"grammaticalFeatures"`
          InflectedForm  string `json:"inflectedForm"`
          Pronunciations []struct {
            AudioFile        string   `json:"audioFile"`
            Dialects         []string `json:"dialects"`
            PhoneticNotation string   `json:"phoneticNotation"`
            PhoneticSpelling string   `json:"phoneticSpelling"`
          } `json:"pronunciations"`
        } `json:"inflections"`
        Pronunciations []struct {
          AudioFile        string   `json:"audioFile"`
          Dialects         []string `json:"dialects"`
          PhoneticNotation string   `json:"phoneticNotation"`
          PhoneticSpelling string   `json:"phoneticSpelling"`
        } `json:"pronunciations"`
        Senses []struct {
          Constructions []struct {
            Text string `json:"text"`
          } `json:"constructions"`
          CrossReferenceMarkers []string `json:"crossReferenceMarkers"`
          CrossReferences       []struct {
            ID   string `json:"id"`
            Text string `json:"text"`
            Type string `json:"type"`
          } `json:"crossReferences"`
          Definitions   []string `json:"definitions"`
          DomainClasses []struct {
            ID   string `json:"id"`
            Text string `json:"text"`
          } `json:"domainClasses"`
          Examples []struct {
            Notes []struct {
              Text string `json:"text"`
              Type string `json:"type"`
            } `json:"notes"`
            Text string `json:"text"`
          } `json:"examples"`
          ID    string `json:"id"`
          Notes []struct {
            Text string `json:"text"`
            Type string `json:"type"`
          } `json:"notes"`
          SemanticClasses []struct {
            ID   string `json:"id"`
            Text string `json:"text"`
          } `json:"semanticClasses"`
          ShortDefinitions []string `json:"shortDefinitions"`
          Subsenses        []struct {
            Definitions   []string `json:"definitions"`
            DomainClasses []struct {
              ID   string `json:"id"`
              Text string `json:"text"`
            } `json:"domainClasses"`
            Examples []struct {
              Notes []struct {
                Text string `json:"text"`
                Type string `json:"type"`
              } `json:"notes"`
              Text string `json:"text"`
            } `json:"examples"`
            ID    string `json:"id"`
            Notes []struct {
              Text string `json:"text"`
              Type string `json:"type"`
            } `json:"notes"`
            Registers []struct {
              ID   string `json:"id"`
              Text string `json:"text"`
            } `json:"registers"`
            SemanticClasses []struct {
              ID   string `json:"id"`
              Text string `json:"text"`
            } `json:"semanticClasses"`
            ShortDefinitions []string `json:"shortDefinitions"`
            Synonyms         []struct {
              Language string `json:"language"`
              Text     string `json:"text"`
            } `json:"synonyms"`
            ThesaurusLinks []struct {
              EntryID string `json:"entry_id"`
              SenseID string `json:"sense_id"`
            } `json:"thesaurusLinks"`
          } `json:"subsenses"`
          Synonyms []struct {
            Language string `json:"language"`
            Text     string `json:"text"`
          } `json:"synonyms"`
          ThesaurusLinks []struct {
            EntryID string `json:"entry_id"`
            SenseID string `json:"sense_id"`
          } `json:"thesaurusLinks"`
        } `json:"senses"`
      } `json:"entries"`
      Language        string `json:"language"`
      LexicalCategory struct {
        ID   string `json:"id"`
        Text string `json:"text"`
      } `json:"lexicalCategory"`
      Phrases []struct {
        ID   string `json:"id"`
        Text string `json:"text"`
      } `json:"phrases"`
      Text string `json:"text"`
    } `json:"lexicalEntries"`
    Type string `json:"type"`
    Word string `json:"word"`
  } `json:"results"`
  Word string `json:"word"`
}
