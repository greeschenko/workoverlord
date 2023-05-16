# WorkOverlord

[DEV] Next-generation personal productivity app powered by mindmaps, secondbrain and AI

# HOW it work

- User run app by first time and set up the SECRETKEY for data enscription 
- User create data map by eding infinity count of primitive data elements (synapses) 
- each synapse can store text, image, video, map, link.. etc. and as many as need nested elements 
- User can add reminding to the any place on the data map
- User can run any system app with reminding 
- all data store only on local device storage
- User can use app on other devises with two-way synchronization between copies

# TODO

- [ ] basic backend
- [ ] backend test
- [ ] encription
- [ ] build in maindmap frontend
- [ ] AI implementation for auto remindings, auto element position etc.
...

# Data structure

```
Maind: [
    Synapse1...,
    Synapse2: {
        size: [2]int
        position: [2]int
        routs: [][2]int
        pointers: [][2]int
        data: string
        synapses: [
            Synapse11...,
            Synapse12...,
            Synapse13...,
            ...
            Synapse1N...,
        ],
        tags: string,
    }
    Synapse3...,
    ...
    SynapseN...,
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
