


- edges: A schema that defines an edge using the edge.To builder owns the relation, unlike using the edge.From builder that gives only a back-reference for the relation (with a different name).# GO-Ent-Edges
  - Relationship
    1. O2O Two Types 'User & Card'
      - In this example, a user has only one credit-card, and a card has only one owner.
  
    2. O2O Same Type 'Node'
      - In this linked-list example, we have a recursive relation named next/prev. Each node in the list can have only one next node
    
    3. O2O Bidirectional 'User'
      - In this user-spouse example, we have a symmetric O2O relation named spouse. Each user can have only one spouse. If user A sets its spouse (using spouse) to B, B can get its spouse using the spouse edge.
      Note that there are no owner/inverse terms in cases of bidirectional edges.
    
    4. O2M Two Types 'User & Pet'
      - In this user-pets example, we have a O2M relation between user and its pets. Each user has many pets, and a pet has one owner. If user A adds a pet B using the pets edge, B can get its owner using the owner edge (the back-reference edge).
      Note that this relation is also a M2O (many-to-one) from the point of view of the Pet schema.
    
    5. O2M Same Type 'Tree'
      - we have a recursive O2M relation between tree's nodes and their children (or their parent).
        Each node in the tree has many children, and has one parent. If node A adds B to its children, B can get its owner using the owner edge.
    
    6. M2M Two Types 'User & Group' <<<<<< New Table >>>>>>
      - we have a M2M relation between groups and their users. Each group has many users, and each user can be joined to many groups.
    
    7. M2M Same Type 'User & Followers'
      - we have a M2M relation between users to their followers. Each user can follow many users, and can have many followers.

    8. 





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