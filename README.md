
- ent.field
  - .Int("") => integer
  - .Float("") => 
  - .Bool("") => boolean
  - .String("") => 
  - .Time("") => sores time
  - .JSON("url", &url.URL{}) => stores a url related to the user
  - .JSON("strings", []string{}) => stores an array of strings
  - .Enum("").Values("on", "off") => stores state of user, Values() method define allowed values
  - .UUID("", uuid.UUID{}) => stores universally unique identifier

- edges
-  A schema that defines an edge using the edge.To builder owns the relation, 
   -  unlike using the edge.From builder that gives only a back-reference for the relation (with a different name).# GO-Ent-Edges
