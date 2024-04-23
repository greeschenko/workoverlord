# WorkOverlord

[DEV] Next-generation personal productivity app powered by mindmaps, secondbrain and AI

![20240418-160914_1920x1080](https://github.com/greeschenko/workoverlord/assets/2754533/3ff2268d-79b9-4de1-8c2f-f852466d6791)

# HOW it work

- User run app by first time and set up the SECRETKEY for data enscription 
- User create data map by eding infinity count of primitive data elements (cells) 
- each cell can store text, image, video, map, link.. etc. and as many as need nested elements 
- User can add reminding to the any place on the data map
- User can run any system app with reminding 
- all data store only on local device storage
- User can use app on other devises with two-way synchronization between copies

# TODO

- [ ] basic backend
- [ ] encription
- [ ] backend test
- [ ] build in maindmap frontend
- [ ] AI implementation for auto remindings, auto element position etc.
...

# Data structure

```
Maind: [                                    //user second brain database object
    Cell1...,                               //data element
    Cell2: {                                //list of child elements
        size: [2]int                        //geometric width and height for frontend
        position: [3]int                    //position on the map, X, Y and Z
        data: string                        //user data
        synapses: [                         //list of connections for an element
            Synapse11...,
            Synapse12{
                points: [][3]int            //connection line ends coordinates X,Y,Z
                size: int                   //line width
                color: string               //line color
                linetype: string            //line type solid | dashed ...
                endtype: string             //line end style none, none | none, arrow | point, point ... etc
            },
            Synapse13...,
            ...
            Synapse1N...,
        ],
        cells: [
            Cell11...,
            Cell12...,
            Cell13...,
            ...
            Cell1N...,
        ],
        tags: string,                       //list of element tags
    }
    Cell3...,
    ...
    CellN...,
]

```

# data enscription strategy

```
     SECRETKEY ---- sha256 encoding user secretkey [32]byte

     local bd = json file with encrypted content ---- AES with 64bit key (SECRETKEY)

     bd init -> promt with user email and password
     -> generate secretkey from password
     -> generate filename sha256(useremail+ random salt + SECRETKEY)
     -> write base json structur

     userhash->taskhash->subtaskhash->subtaskhash->...
        |          |                     /
        |          |                     \____ encrypted content ---- AES with 64bit key (SECRETKEY)
        |          |
        |          |
        |          \___sha256(parent hash + SECRETKEY)
        |
        |
       sha256(useremail + SECRETKEY)
```
