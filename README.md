# ytsearch

### Search for YouTube videos, channels & playlists. Without YouTube Data API.

#### A Go port of [alexmercerind](https://github.com/alexmercerind)'s [youtube-search-python](https://github.com/alexmercerind/youtube-search-python).

## Installing

```bash
go get github.com/raitonoberu/ytsearch
```

Please note that the API is not yet final and may change in the future.

## Usage

### Search for videos only

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/raitonoberu/ytsearch"
)

func main() {
    search := ytsearch.VideoSearch("ncs")
    result, err := search.Next()
    if err != nil {
        panic(err)
    }

    jsonstr, _ := json.Marshal(result.Videos[0])
    fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
  "id":"-ObdvMkCKws",
  "title":"Top 20 Most Popular Songs by NCS | Best of NCS | Most Viewed Songs",
  "publishedTime":"2 years ago",
  "duration":4367,
  "viewCount":28009374,
  "thumbnails":[
    {
      "url":"https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCXso3cln-2aZPAiSNW15njQ-xchw",
      "height":202,
      "width":360
    },
    {
      "url":"https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCYGQ4cDZ49e0mlnD2Uj_oAI03wzQ",
      "height":404,
      "width":720
    }
  ],
  "richThumbnail":{
    "url":"",
    "height":0,
    "width":0
  },
  "description":"Top NCS full ☞ https://youtu.be/HPhHr6h4Qjc ☞ Follow Magic Sound : Facebook → https://goo.gl/fWA96C Subscribe ...",
  "channel":{
    "id":"UCoDZIZuadPBixSPFR7jAq2A",
    "title":"Music Store",
    "thumbnails":[
      {
        "url":"https://yt3.ggpht.com/ytc/AKedOLSNH-mOdxWiiqLGpFJAL6Zq87MvN7w9I42Hr9a_=s68-c-k-c0x00ffffff-no-rj",
        "height":68,
        "width":68
      }
    ],
    "url":"https://www.youtube.com/channel/UCoDZIZuadPBixSPFR7jAq2A"
  },
  "url":"https://www.youtube.com/watch?v=-ObdvMkCKws"
}
```
</details>

### Search for channels only

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/raitonoberu/ytsearch"
)

func main() {
    search := ytsearch.ChannelSearch("ncs")
    result, err := search.Next()
    if err != nil {
        panic(err)
    }

    jsonstr, _ := json.Marshal(result.Channels[0])
    fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
  "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
  "title":"NoCopyrightSounds",
  "thumbnails":[
    {
      "url":"http://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s88-c-k-c0x00ffffff-no-rj-mo",
      "height":88,
      "width":88
    },
    {
      "url":"http://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s176-c-k-c0x00ffffff-no-rj-mo",
      "height":176,
      "width":176
    }
  ],
  "videoCount":966,
  "description":"NoCopyrightSounds is a copyright free / stream safe record label, providing free to use music to the content creator community.",
  "subscribers":"31M subscribers",
  "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
}
```
</details>

### Search for playlists only

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/raitonoberu/ytsearch"
)

func main() {
    search := ytsearch.PlaylistSearch("ncs")
    result, err := search.Next()
    if err != nil {
        panic(err)
    }

    jsonstr, _ := json.Marshal(result.Playlists[0])
    fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
  "id":"PLzkuLC6Yvumv_Rd5apfPRWEcjf9b1JRnq",
  "title":"NCS Playlist Best Songs",
  "videoCount":792,
  "channel":{
    "id":"UC7o5WwsbRAdfZwDe2brud7g",
    "title":"NCS",
    "thumbnails":null,
    "url":"https://www.youtube.com/channel/UC7o5WwsbRAdfZwDe2brud7g"
  },
  "thumbnails":[
    {
      "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\u0026rs=AOn4CLAD6Vv5Z0roTmEwnNF_NyBhfJKN-w",
      "height":94,
      "width":168
    },
    {
      "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEWCMQBEG5IWvKriqkDCQgBFQAAiEIYAQ==\u0026rs=AOn4CLDxRqXhHvPhehwYkIkGhQkGx9OswQ",
      "height":110,
      "width":196
    },
    {
      "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=\u0026rs=AOn4CLCcakX_wGhCA1InSGO01CTTVn0AdQ",
      "height":138,
      "width":246
    },
    {
      "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=\u0026rs=AOn4CLCcWVWWwwNAxJ_F5NOVwGXepJxsQA",
      "height":188,
      "width":336
    }
  ],
  "url":"https://www.youtube.com/playlist?list=PLzkuLC6Yvumv_Rd5apfPRWEcjf9b1JRnq"
}
```
</details>

### Search for everything

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/raitonoberu/ytsearch"
)

func main() {
    search := ytsearch.Search("ncs")
    result, err := search.Next()
    if err != nil {
        panic(err)
    }

    jsonstr, _ := json.Marshal(result)
    fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
  "estimatedResults":29030917,
  "videos":[
    {
      "id":"-ObdvMkCKws",
      "title":"Top 20 Most Popular Songs by NCS | Best of NCS | Most Viewed Songs",
      "publishedTime":"2 years ago",
      "duration":4367,
      "viewCount":28009499,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCXso3cln-2aZPAiSNW15njQ-xchw",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCYGQ4cDZ49e0mlnD2Uj_oAI03wzQ",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"",
        "height":0,
        "width":0
      },
      "description":"Top NCS full ☞ https://youtu.be/HPhHr6h4Qjc ☞ Follow Magic Sound : Facebook → https://goo.gl/fWA96C Subscribe ...",
      "channel":{
        "id":"UCoDZIZuadPBixSPFR7jAq2A",
        "title":"Music Store",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLSNH-mOdxWiiqLGpFJAL6Zq87MvN7w9I42Hr9a_=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UCoDZIZuadPBixSPFR7jAq2A"
      },
      "url":"https://www.youtube.com/watch?v=-ObdvMkCKws"
    },
    {
      "id":"JNl1_hRwpXE",
      "title":"NCS: 30 Million Subscriber Mix",
      "publishedTime":"4 months ago",
      "duration":5726,
      "viewCount":5150157,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/JNl1_hRwpXE/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDwkR7V56OMD5UY9uwyr1ImIAu97Q",
          "height":270,
          "width":480
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/JNl1_hRwpXE/mqdefault_6s.webp?du=3000\u0026sqp=CKC3uIkG\u0026rs=AOn4CLBnNKmtUEQU-qMr1w-mW1CB-tpDwg",
        "height":180,
        "width":320
      },
      "description":"Best of NoCopyrightSounds 2021 (30 Million Subscriber Mix) WE DID IT!! ❤️ Thank you SO much for all the support you guys ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=JNl1_hRwpXE"
    },
    {
      "id":"mt2hJr4ulLU",
      "title":"🔥 Top 50 NoCopyRightSounds | Best of NCS | Most viewed ! Gaming Music | The Best of All Time | 2020",
      "publishedTime":"1 year ago",
      "duration":11286,
      "viewCount":2847126,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/mt2hJr4ulLU/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBRHb79Nc1CJFFDdRgecoEyrgrI5A",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/mt2hJr4ulLU/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLD_6CNVzwjKoj6XBpCeU7bQr_2G9A",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/mt2hJr4ulLU/mqdefault_6s.webp?du=3000\u0026sqp=CNPtuIkG\u0026rs=AOn4CLBuLQTTkFmTI3PYXxHlDWZAINzPQg",
        "height":180,
        "width":320
      },
      "description":"Best of NCS | The Best of All Time Track list: 01. 0:00 Alan Walker - Fade 02. 4:20 Alan Walker - Spectre 03. 8:07 Different Heaven ...",
      "channel":{
        "id":"UCohd6FLJfy3JMmYWJcn4yYA",
        "title":"Phan Trieu",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQ8mo8_0tF9IxgWhyjiDO3ImqA5jTH3l5hxwtU_aA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UCohd6FLJfy3JMmYWJcn4yYA"
      },
      "url":"https://www.youtube.com/watch?v=mt2hJr4ulLU"
    },
    {
      "id":"eUR-DrTGNws",
      "title":"Top 10 Most Popular Songs by NCS | Episode 1",
      "publishedTime":"9 months ago",
      "duration":2404,
      "viewCount":16403579,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/eUR-DrTGNws/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLD0A-CGEUwsgAWfryyOArKag_GUXw",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/eUR-DrTGNws/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCqaT6eE5C0eNJNgDkCD9vrE7p5IQ",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/eUR-DrTGNws/mqdefault_6s.webp?du=3000\u0026sqp=CJnSuIkG\u0026rs=AOn4CLAr2vubXjYR1iJ7rB_lJIp_0iY1YA",
        "height":180,
        "width":320
      },
      "description":"0:00 - 1. Alan Walker - Fade 4:20 - 2. Cartoon - On \u0026 On 7:48 - 3. Alan Walker - Spectre 11:34 - 4. DEAF KEV - Invincible 16:07 - 5.",
      "channel":{
        "id":"UC0raGe2owvmadTf85xL_ytw",
        "title":"Limon NCS",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLTbr1Jn_Ey7zKqrnmhQ-LVs6R0O6evUHLHQm2mM=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC0raGe2owvmadTf85xL_ytw"
      },
      "url":"https://www.youtube.com/watch?v=eUR-DrTGNws"
    },
    {
      "id":"triXo_xCqms",
      "title":"TOP 100 NoCopyrightSounds | Best Of NCS, 6H NoCopyRightSounds | TOP 100 NCS, The Best of all time",
      "publishedTime":"3 years ago",
      "duration":22664,
      "viewCount":7343922,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/triXo_xCqms/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCCdwpnnd68xW-Yg-9oz-L3kXWutg",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/triXo_xCqms/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLD5VNhq765byHhvrqCMm7vh-cx4dQ",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"",
        "height":0,
        "width":0
      },
      "description":"TOP 100 NoCopyrightSounds | Best Of NCS | 6H NoCopyRightSounds | The Best of all time NCS Song: Best of all time Please ...",
      "channel":{
        "id":"UCDsbBjIl0lYZB4IokDLyWIQ",
        "title":"Musicbot",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLSpQ7TO0rX8jEfIp0hd3sYODzN800Us7gGfbKmvrQ=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UCDsbBjIl0lYZB4IokDLyWIQ"
      },
      "url":"https://www.youtube.com/watch?v=triXo_xCqms"
    },
    {
      "id":"bAupUaalXLE",
      "title":"게임할때 듣기 좋은 NCS 노래 모음[롤 할 때 들으면 캐리 가능]",
      "publishedTime":"11 months ago",
      "duration":7180,
      "viewCount":2812770,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/bAupUaalXLE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDKB9hoBGIdHlQNY9bQUeyv0lZHFw",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/bAupUaalXLE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAsdBJEiQxPAPRzUeO0RQYCe698OQ",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/bAupUaalXLE/mqdefault_6s.webp?du=3000\u0026sqp=CJfNuIkG\u0026rs=AOn4CLBSHMtelGMS287RBIXwYpHOnd07gQ",
        "height":180,
        "width":320
      },
      "description":"게임할때듣기좋은노래 #NCS노래 #롤노래 #게임노래 #매드무비브금 [소소한 후원] 구독과 좋아요는 큰 힘이 됩니다. [BGM 정보] ...",
      "channel":{
        "id":"UC1t-vzAacQ4OiMLt5h7Flug",
        "title":"찰리의 롤튜브",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQ2VDHsi2CZ5AvHGvieaad2TKehcKltK1G7nJ2u7w=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC1t-vzAacQ4OiMLt5h7Flug"
      },
      "url":"https://www.youtube.com/watch?v=bAupUaalXLE"
    },
    {
      "id":"pSFVZ1OjQpk",
      "title":"🔥Fantastic Mix For Gaming: Top 30 Songs ♫ Best NCS Gaming Music ♫ Best Of EDM 2021",
      "publishedTime":"1 day ago",
      "duration":10800,
      "viewCount":7196,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/pSFVZ1OjQpk/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBTdHZP_K9YV1d5Tw17M_-XcyTwXA",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/pSFVZ1OjQpk/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBqJLdT7w8tqaklkRojX3w5f3yTMg",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/pSFVZ1OjQpk/mqdefault_6s.webp?du=3000\u0026sqp=CLzKuIkG\u0026rs=AOn4CLC5LcPemZsSB-UGzyGxoBeI1V--3g",
        "height":180,
        "width":320
      },
      "description":"Welcome to Freeme NCS☆ ------- ▻Fantastic Mix For Gaming: Top 30 Songs ♫ Best NCS Gaming Music ♫ Best Of EDM 2021 ...",
      "channel":{
        "id":"UCG0SzK_t4-Ylf1yZq9Xmi_g",
        "title":"Freeme NCS Music",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLSHgpuCBptK0p-AbLOrcb-Utpxr4zDJvPAddlf2=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UCG0SzK_t4-Ylf1yZq9Xmi_g"
      },
      "url":"https://www.youtube.com/watch?v=pSFVZ1OjQpk"
    },
    {
      "id":"yJg-Y5byMMw",
      "title":"Warriyo - Mortals (feat. Laura Brehm) [NCS Release]",
      "publishedTime":"4 years ago",
      "duration":230,
      "viewCount":190268510,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/yJg-Y5byMMw/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAmKIY6QpyYwke2y20r8AbLNnEQjw",
          "height":270,
          "width":480
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/yJg-Y5byMMw/mqdefault_6s.webp?du=3000\u0026sqp=CKDYuIkG\u0026rs=AOn4CLCz0fotR3U6wUt8bbe6oPiN0i-4TA",
        "height":180,
        "width":320
      },
      "description":"NCS: Music Without Limitations NCS Spotify: http://spoti.fi/NCS Free Download / Stream: http://ncs.io/mortals Connect with NCS: ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=yJg-Y5byMMw"
    },
    {
      "id":"1dVPjcmEUsk",
      "title":"🔥Amazing EDM Music 2021 Mix ♫ Top 30 NCS Gaming Music - Vocal Mix ♫ Best Of EDM 2021",
      "publishedTime":"2 days ago",
      "duration":10800,
      "viewCount":8821,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/1dVPjcmEUsk/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCK2U7Gud26t74J06boKR4qIj2wMA",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/1dVPjcmEUsk/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLC5jSat0UWzag1vkQPboTeyu7sQ8g",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/1dVPjcmEUsk/mqdefault_6s.webp?du=3000\u0026sqp=CNzkuIkG\u0026rs=AOn4CLAsEg8mmyDH6fo9DZicaSsY2J8AEg",
        "height":180,
        "width":320
      },
      "description":"Welcome to Freeme NCS☆ ------- ▻Amazing EDM Music 2021 Mix ♫ Top 30 NCS Gaming Music - Vocal Mix ♫ Best Of EDM 2021 ...",
      "channel":{
        "id":"UCG0SzK_t4-Ylf1yZq9Xmi_g",
        "title":"Freeme NCS Music",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLSHgpuCBptK0p-AbLOrcb-Utpxr4zDJvPAddlf2=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UCG0SzK_t4-Ylf1yZq9Xmi_g"
      },
      "url":"https://www.youtube.com/watch?v=1dVPjcmEUsk"
    },
    {
      "id":"i6WvvstJ-Dw",
      "title":"N3WPORT \u0026 Meggie York - Runaway [NCS Release]",
      "publishedTime":"5 months ago",
      "duration":213,
      "viewCount":1207946,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/i6WvvstJ-Dw/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDpCog_jkbHsoDGQppOqdSTAj59Gw",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/i6WvvstJ-Dw/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAfSx0SMw7gd5QTVa61Lnq3MvJ6mQ",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"",
        "height":0,
        "width":0
      },
      "description":"NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=i6WvvstJ-Dw"
    },
    {
      "id":"ABuNwLP-z9o",
      "title":"🔥 Top 50 NoCopyRightSounds | Best of NCS | Most viewed ! Gaming Music | The Best of All Time | 2021",
      "publishedTime":"2 years ago",
      "duration":10822,
      "viewCount":9946257,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/ABuNwLP-z9o/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLC0rl3DGDwIvQs3ADcVp40EyQmhDw",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/ABuNwLP-z9o/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBUWW_Gz6kYwRcNJ9yudEipM2idwQ",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/ABuNwLP-z9o/mqdefault_6s.webp?du=3000\u0026sqp=CIriuIkG\u0026rs=AOn4CLAXcior3KjUzmwjd-dsgPQm9P3HgA",
        "height":180,
        "width":320
      },
      "description":"Best of NCS | The Best of All Time Track list: 01. 0:00 Alan Walker - Fade 02. 4:20 Alan Walker - Spectre 03. 8:07 Different Heaven ...",
      "channel":{
        "id":"UCG0SzK_t4-Ylf1yZq9Xmi_g",
        "title":"Freeme NCS Music",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLSHgpuCBptK0p-AbLOrcb-Utpxr4zDJvPAddlf2=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UCG0SzK_t4-Ylf1yZq9Xmi_g"
      },
      "url":"https://www.youtube.com/watch?v=ABuNwLP-z9o"
    },
    {
      "id":"p7ZsBPK656s",
      "title":"Disfigure - Blank [NCS Release]",
      "publishedTime":"8 years ago",
      "duration":210,
      "viewCount":212344698,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/p7ZsBPK656s/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBv_yXtB8uYVBTuW1vvzL4nIzWR0w",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/p7ZsBPK656s/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAeVk073JnAfD9Nd6sRe8zxlkDnyA",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/p7ZsBPK656s/mqdefault_6s.webp?du=3000\u0026sqp=CI3RuIkG\u0026rs=AOn4CLD6fJZSAQOTkkp-9ekxLGKRtjfW8A",
        "height":180,
        "width":320
      },
      "description":"NCS: Music Without Limitations NCS Spotify: http://spoti.fi/NCS Download / Stream: http://ncs.io/blank ▽ Connect with NCS ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=p7ZsBPK656s"
    },
    {
      "id":"QglaLzo_aPk",
      "title":"Julius Dreisig \u0026 Zeus X Crona - Invisible [NCS Release]",
      "publishedTime":"2 years ago",
      "duration":201,
      "viewCount":113693824,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/QglaLzo_aPk/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCo0PmNWWnwg_Y_H540C-GnyjdsOw",
          "height":270,
          "width":480
        }
      ],
      "richThumbnail":{
        "url":"",
        "height":0,
        "width":0
      },
      "description":"- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your description: ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=QglaLzo_aPk"
    },
    {
      "id":"K4DyBUG242c",
      "title":"Cartoon - On \u0026 On (feat. Daniel Levi) [NCS Release]",
      "publishedTime":"6 years ago",
      "duration":208,
      "viewCount":428959970,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/K4DyBUG242c/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCVDGZSQEPvPuF1QyELmT2FV6vdaQ",
          "height":270,
          "width":480
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/K4DyBUG242c/mqdefault_6s.webp?du=3000\u0026sqp=CJm2uIkG\u0026rs=AOn4CLD-AtNaB6yoF5UatJmWs_qYtDwuRw",
        "height":180,
        "width":320
      },
      "description":"Lyrics: Cartoon - On \u0026 On feat. Daniel Levi [Verse 1] Hold me close 'til I get up Time is barely on our side I don't wanna waste ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=K4DyBUG242c"
    },
    {
      "id":"S19UcWdOA-I",
      "title":"Lost Sky - Fearless pt.II (feat. Chris Linton) [NCS Release]",
      "publishedTime":"3 years ago",
      "duration":194,
      "viewCount":133478422,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/S19UcWdOA-I/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDMiju4cvt7r4RId6P3lLQYUepbpg",
          "height":270,
          "width":480
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/S19UcWdOA-I/mqdefault_6s.webp?du=3000\u0026sqp=CPrauIkG\u0026rs=AOn4CLCnkXrV1V4BrPPcsN9JMb7CUEJGsA",
        "height":180,
        "width":320
      },
      "description":"- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=S19UcWdOA-I"
    },
    {
      "id":"EnuGdwI0W1g",
      "title":"Elektronomia \u0026 RUD - Memory [NCS Release]",
      "publishedTime":"4 months ago",
      "duration":255,
      "viewCount":2174680,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/EnuGdwI0W1g/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAjOq_xXQcN9w6DrOdP9S2_OBgNYg",
          "height":202,
          "width":360
        },
        {
          "url":"https://i.ytimg.com/vi/EnuGdwI0W1g/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBycqhUQMk7wCJhP0kBbSeX0yQyVg",
          "height":404,
          "width":720
        }
      ],
      "richThumbnail":{
        "url":"https://i.ytimg.com/an_webp/EnuGdwI0W1g/mqdefault_6s.webp?du=3000\u0026sqp=CLi9uIkG\u0026rs=AOn4CLCVr-PcpA_jhntfZvpX7DzMIvI0hg",
        "height":180,
        "width":320
      },
      "description":"NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=EnuGdwI0W1g"
    },
    {
      "id":"TW9d8vYrVFQ",
      "title":"Elektronomia - Sky High [NCS Release]",
      "publishedTime":"4 years ago",
      "duration":238,
      "viewCount":200001954,
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/TW9d8vYrVFQ/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDzHv6boRaBnEoZWSNFdZRGVdUmJw",
          "height":270,
          "width":480
        }
      ],
      "richThumbnail":{
        "url":"",
        "height":0,
        "width":0
      },
      "description":"NCS: Music Without Limitations NCS Spotify: http://spoti.fi/NCS Free Download / Stream: http://ncs.io/skyhigh Connect with NCS: ...",
      "channel":{
        "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title":"NoCopyrightSounds",
        "thumbnails":[
          {
            "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
            "height":68,
            "width":68
          }
        ],
        "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "url":"https://www.youtube.com/watch?v=TW9d8vYrVFQ"
    }
  ],
  "channels":[
    {
      "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
      "title":"NoCopyrightSounds",
      "thumbnails":[
        {
          "url":"http://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s88-c-k-c0x00ffffff-no-rj-mo",
          "height":88,
          "width":88
        },
        {
          "url":"http://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s176-c-k-c0x00ffffff-no-rj-mo",
          "height":176,
          "width":176
        }
      ],
      "videoCount":966,
      "description":"NoCopyrightSounds is a copyright free / stream safe record label, providing free to use music to the content creator community.",
      "subscribers":"31M subscribers",
      "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
    }
  ],
  "playlists":[
    {
      "id":"PLzkuLC6Yvumv_Rd5apfPRWEcjf9b1JRnq",
      "title":"NCS Playlist Best Songs",
      "videoCount":792,
      "channel":{
        "id":"UC7o5WwsbRAdfZwDe2brud7g",
        "title":"NCS",
        "thumbnails":null,
        "url":"https://www.youtube.com/channel/UC7o5WwsbRAdfZwDe2brud7g"
      },
      "thumbnails":[
        {
          "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\u0026rs=AOn4CLAD6Vv5Z0roTmEwnNF_NyBhfJKN-w",
          "height":94,
          "width":168
        },
        {
          "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEWCMQBEG5IWvKriqkDCQgBFQAAiEIYAQ==\u0026rs=AOn4CLDxRqXhHvPhehwYkIkGhQkGx9OswQ",
          "height":110,
          "width":196
        },
        {
          "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=\u0026rs=AOn4CLCcakX_wGhCA1InSGO01CTTVn0AdQ",
          "height":138,
          "width":246
        },
        {
          "url":"http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=\u0026rs=AOn4CLCcWVWWwwNAxJ_F5NOVwGXepJxsQA",
          "height":188,
          "width":336
        }
      ],
      "url":"https://www.youtube.com/playlist?list=PLzkuLC6Yvumv_Rd5apfPRWEcjf9b1JRnq"
    }
  ],
  "shelves":[
    {
      "title":"Latest from NoCopyrightSounds",
      "items":[
        {
          "id":"lj62iuaKAhU",
          "title":"JPB - High (feat. Aleesia) [NCS10 Release]",
          "publishedTime":"2 days ago",
          "duration":153,
          "viewCount":430203,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/lj62iuaKAhU/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBuUh1qTqRYcCSIW-F9A1x6zfZMWA",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/lj62iuaKAhU/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLC3BLjJWMPagc7oY07PaK6QQ0RpRw",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/lj62iuaKAhU/mqdefault_6s.webp?du=3000\u0026sqp=CJ3quIkG\u0026rs=AOn4CLAKM6UmdGCVN5e1tzfOmqxPwPuj9w",
            "height":180,
            "width":320
          },
          "description":"NCS (NoCopyrightSounds): Empowering Creators through Copyright / Royalty Free Music Follow us on Spotify: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=lj62iuaKAhU"
        },
        {
          "id":"w8xmc7TrX-c",
          "title":"Culture Code - You \u0026 I (feat. Alexis Donn) [NCS10 Release]",
          "publishedTime":"4 days ago",
          "duration":215,
          "viewCount":421830,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/w8xmc7TrX-c/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDsDh8GRXdnP8LPtbHmUIOFDRalkQ",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/w8xmc7TrX-c/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDvN2JM9fp2YQijQO4eVMMXgAqfrw",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/w8xmc7TrX-c/mqdefault_6s.webp?du=3000\u0026sqp=COjbuIkG\u0026rs=AOn4CLBThIxKfJlbog-Hj68DmVQ_7pXh9A",
            "height":180,
            "width":320
          },
          "description":"NCS (NoCopyrightSounds): Empowering Creators through Copyright / Royalty Free Music Follow us on Spotify: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=w8xmc7TrX-c"
        },
        {
          "id":"AssxzJCFl24",
          "title":"THYKIER - Station 2 [NCS10 Release]",
          "publishedTime":"6 days ago",
          "duration":190,
          "viewCount":390221,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/AssxzJCFl24/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLACwfY-7sScLmfpx3JgznWKgmEQjA",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/AssxzJCFl24/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAjR7uWRxkGyH7AjWWYLOltaDeqoA",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/AssxzJCFl24/mqdefault_6s.webp?du=3000\u0026sqp=CNLmuIkG\u0026rs=AOn4CLAw1kgIPH2t4IgXa08sIAiz3g9Tog",
            "height":180,
            "width":320
          },
          "description":"- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your description: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=AssxzJCFl24"
        },
        {
          "id":"yH88qRmgkGI",
          "title":"Netrum \u0026 Halvorsen - Phoenix [NCS10 Release]",
          "publishedTime":"1 week ago",
          "duration":238,
          "viewCount":698643,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/yH88qRmgkGI/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCZiabJw-fQ9XsTO1CarVpWgj9nFA",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/yH88qRmgkGI/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLA3JGe42Lndpakjoboi-8W4cgZW4w",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/yH88qRmgkGI/mqdefault_6s.webp?du=3000\u0026sqp=CJa7uIkG\u0026rs=AOn4CLA4atblDiuTtwbZoJCQuiTpZFurOg",
            "height":180,
            "width":320
          },
          "description":"- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your description: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=yH88qRmgkGI"
        },
        {
          "id":"QVKwpSRr3oY",
          "title":"Razihel - A Song About You [NCS10 Release]",
          "publishedTime":"1 week ago",
          "duration":148,
          "viewCount":480526,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/QVKwpSRr3oY/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAr2MxpXBEu9wnhaYX0QaQ9N461eg",
              "height":270,
              "width":480
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/QVKwpSRr3oY/mqdefault_6s.webp?du=3000\u0026sqp=CNbSuIkG\u0026rs=AOn4CLDeFHRjT7lq8FaNikDmFlo9RHf9-g",
            "height":180,
            "width":320
          },
          "description":"NCS (NoCopyrightSounds): Empowering Creators through Copyright / Royalty Free Music Follow us on Spotify: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=QVKwpSRr3oY"
        },
        {
          "id":"-mMmOKHzuWc",
          "title":"Jim Yosef, Electro-Light, Anna Yvette, Deaf Kev \u0026 Tobu - Destiny [NCS10 Release]",
          "publishedTime":"1 week ago",
          "duration":229,
          "viewCount":698759,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/-mMmOKHzuWc/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDgphZBxL1le0woqknbRCrAO3z-dg",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/-mMmOKHzuWc/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAfGEZ2uNWnpTHeDyDLv_6XbA0wIQ",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/-mMmOKHzuWc/mqdefault_6s.webp?du=3000\u0026sqp=CMHMuIkG\u0026rs=AOn4CLDPe6Jw8J9Bxqxvv5JsIOQtJqxFSQ",
            "height":180,
            "width":320
          },
          "description":"NCS (NoCopyrightSounds): Empowering Creators through Copyright / Royalty Free Music Follow us on Spotify: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=-mMmOKHzuWc"
        },
        {
          "id":"8-OtWif8Z4A",
          "title":"The Creation of NCS",
          "publishedTime":"2 weeks ago",
          "duration":73,
          "viewCount":664772,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/8-OtWif8Z4A/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLC0_VDfh_Ownc5cOCmikSh492fLmw",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/8-OtWif8Z4A/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCB14Mvlz4RkWi0fC47fX3UpZNytQ",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/8-OtWif8Z4A/mqdefault_6s.webp?du=3000\u0026sqp=CODcuIkG\u0026rs=AOn4CLDqwccfIxVGovk8DghL9m64aG6UnA",
            "height":180,
            "width":320
          },
          "description":"10 years ago today, NoCopyrightSounds was born. I sat down at my computer, opened YouTube \u0026 created this channel. I wanted ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=8-OtWif8Z4A"
        },
        {
          "id":"Ueg-SUFGb4M",
          "title":"Alex Skrindo, Severin \u0026 Like Lions - Heart [NCS Release]",
          "publishedTime":"2 weeks ago",
          "duration":163,
          "viewCount":683745,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/Ueg-SUFGb4M/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBVp94D7IKJJRN8oXfgGbg3U-y1Qg",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/Ueg-SUFGb4M/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAJhmNKyYlp-pCgZe_E94OjqMkUQQ",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/Ueg-SUFGb4M/mqdefault_6s.webp?du=3000\u0026sqp=CIXVuIkG\u0026rs=AOn4CLBJSr_ImBMLX5_6G-dIjo9BTznIOg",
            "height":180,
            "width":320
          },
          "description":"NCS (NoCopyrightSounds): Empowering Creators through Copyright / Royalty Free Music Follow us on Spotify: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=Ueg-SUFGb4M"
        },
        {
          "id":"25-ma4FUXuY",
          "title":"NATSUMI - Take Me Away [NCS Release]",
          "publishedTime":"2 weeks ago",
          "duration":174,
          "viewCount":588567,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/25-ma4FUXuY/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLC4339JFdGHW9xs6XLQxRdWqDWjDg",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/25-ma4FUXuY/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLB28-VgeKVexjhzAxgDZ2XMgI8_Gg",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/25-ma4FUXuY/mqdefault_6s.webp?du=3000\u0026sqp=CKrjuIkG\u0026rs=AOn4CLBDz_VyC5PFPnyXQaHpCu-YIsXDbw",
            "height":180,
            "width":320
          },
          "description":"- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your description: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=25-ma4FUXuY"
        },
        {
          "id":"-aSuZsRsOys",
          "title":"Simbai \u0026 Frizzy The Streetz - Crazy [NCS Release]",
          "publishedTime":"3 weeks ago",
          "duration":181,
          "viewCount":483746,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/-aSuZsRsOys/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLClXCqrBkGJvx4gusfCyaMLUeLpHA",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/-aSuZsRsOys/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDOgVwBQgbE9XCcMBSfbmp4TAmvRQ",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/-aSuZsRsOys/mqdefault_6s.webp?du=3000\u0026sqp=CJHkuIkG\u0026rs=AOn4CLC1Hn0cZ4jWkB8kqaoEsELYEwkimA",
            "height":180,
            "width":320
          },
          "description":"NCS (NoCopyrightSounds): Empowering Creators through Copyright / Royalty Free Music Follow us on Spotify: ...",
          "channel":{
            "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title":"NoCopyrightSounds",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "url":"https://www.youtube.com/watch?v=-aSuZsRsOys"
        }
      ]
    },
    {
      "title":"People also watched",
      "items":[
        {
          "id":"QcryKJS3fAY",
          "title":"글자 생방 부캐[나로] 시너지를 위해 배메에 집 보증금 떄려박는다고 오셨습니다..  메이플",
          "publishedTime":"",
          "duration":0,
          "viewCount":0,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/QcryKJS3fAY/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCST9BndL6U1pKFSPnsoLhvjUwXPw",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/QcryKJS3fAY/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAB44KwXN_hrZjRJl5IeRuSrkqf6Q",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"",
            "height":0,
            "width":0
          },
          "description":"항상 영상 사랑해주셔서 감사합니다! 세글자 아프리카TV : http://afreeca.com/skswhdkgo 세글자 인스타 ...",
          "channel":{
            "id":"UCb5NLtXAsTBrmaZVhyFa-Wg",
            "title":"��자네 YouTube",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQJBbz7waUv7RLKoyhb_v6B5T4IdfWIJhauyCxCRw=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCb5NLtXAsTBrmaZVhyFa-Wg"
          },
          "url":"https://www.youtube.com/watch?v=QcryKJS3fAY"
        },
        {
          "id":"IvVgQHZtpQE",
          "title":"Best of Elektronomia 2018 | Top 20 Songs of Elektronomia",
          "publishedTime":"3 years ago",
          "duration":4757,
          "viewCount":2655512,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/IvVgQHZtpQE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDMuVnG1n3vFCydcS25IoGFeATqiA",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/IvVgQHZtpQE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLA3EB6Wp6JT-cDSieyTiTqWcFTCVA",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"",
            "height":0,
            "width":0
          },
          "description":"Elektronomia Top 20 | Best of Elektronomia 2018 Please help Musicbot reach 100k subscriber by join at ...",
          "channel":{
            "id":"UCDsbBjIl0lYZB4IokDLyWIQ",
            "title":"Musicbot",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLSpQ7TO0rX8jEfIp0hd3sYODzN800Us7gGfbKmvrQ=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCDsbBjIl0lYZB4IokDLyWIQ"
          },
          "url":"https://www.youtube.com/watch?v=IvVgQHZtpQE"
        },
        {
          "id":"D-FIfQ5bINo",
          "title":"Top 50 NoCopyRightSounds | Best of NCS | The Best of All Time | Mp3 download | 2021",
          "publishedTime":"1 year ago",
          "duration":11025,
          "viewCount":1352635,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/D-FIfQ5bINo/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCHzw5bIbSDe0waGIzmja8BCuA_RQ",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/D-FIfQ5bINo/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCgrgm0PNy3Zarm2OekUo6Vj06Usw",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"",
            "height":0,
            "width":0
          },
          "description":"* None of these images, music \u0026 video clips were created/owned by us. * This video is purely fan-made, if you (owners) want to ...",
          "channel":{
            "id":"UCAirLxKHLThZrEtJGI_J9kw",
            "title":"Đinh Giang",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQkIXdcR2kGsoJnYKqmAafzdjff07B-UqSynzLByj0=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCAirLxKHLThZrEtJGI_J9kw"
          },
          "url":"https://www.youtube.com/watch?v=D-FIfQ5bINo"
        },
        {
          "id":"A8DYb2MwbCM",
          "title":"【ダイエット】８月最後のリングフィット",
          "publishedTime":"",
          "duration":0,
          "viewCount":0,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/A8DYb2MwbCM/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAuZdsy4SW1brKTDCHruxwO1q6qJg",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/A8DYb2MwbCM/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBPHHo_W-nl81YvZwA4kRQDI5djnQ",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"",
            "height":0,
            "width":0
          },
          "description":"コメントで喧嘩するような人にはならないで メンバー→https://www.youtube.com/channel/UC3NHhO0wBaZ5OVxaabcDKrw/join ...",
          "channel":{
            "id":"UC3NHhO0wBaZ5OVxaabcDKrw",
            "title":"ありけん【クラスに一人はいそうなデブCh】",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/MnZvzwej22BJ5LOCgc0BNomO8jHH9k7yuCYlzXeFWpjAXxjT63nGpRDUaXRs5M8FKklWznGN2A=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UC3NHhO0wBaZ5OVxaabcDKrw"
          },
          "url":"https://www.youtube.com/watch?v=A8DYb2MwbCM"
        },
        {
          "id":"S8SjwKMGswM",
          "title":"Warriyo - Mortals (feat. Laura Brehm) 【1 HOUR】",
          "publishedTime":"4 years ago",
          "duration":3680,
          "viewCount":5217247,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/S8SjwKMGswM/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDUCtQ-VQTv3CMF1qcMsbV6at7Ffg",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/S8SjwKMGswM/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLAfD6rbM1Z9--C7xh62kEUeqZyBtA",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/S8SjwKMGswM/mqdefault_6s.webp?du=3000\u0026sqp=CIDPuIkG\u0026rs=AOn4CLA4afyZDLOwGN0zH9sw77Iawa3A7g",
            "height":180,
            "width":320
          },
          "description":"▭▭▭▭▭▭▭▭▭▭▭▭▭▭▭▭▭▭ NoCopyrightSounds is the record label that connects content creators with the finest ...",
          "channel":{
            "id":"UCs5wn_9Kp-29s0lKUkya-uQ",
            "title":"1 Hour Music",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLQO6wrhnYpnBJbw9PSvmW0TaaIm9eqephPLygJ-=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCs5wn_9Kp-29s0lKUkya-uQ"
          },
          "url":"https://www.youtube.com/watch?v=S8SjwKMGswM"
        },
        {
          "id":"6vztdGgJaAw",
          "title":"Alan Walker - Spectre [1 Hour Version] - NCS Release [Free Download]",
          "publishedTime":"5 years ago",
          "duration":3613,
          "viewCount":5065882,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/6vztdGgJaAw/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBafKTABuUAVWYXZgMkYooCSenwOA",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/6vztdGgJaAw/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLC__4mvnkoJDjsiy18J4rWjeAMecg",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/6vztdGgJaAw/mqdefault_6s.webp?du=3000\u0026sqp=CKyyuIkG\u0026rs=AOn4CLA6oln8TEGMa30W4lRO0-bBtyTzWA",
            "height":180,
            "width":320
          },
          "description":"NoCopyrightSounds, music without limitations. ▽ Follow with NCS Twitch http://twitch.tv/nocopyrightsounds Spotify ...",
          "channel":{
            "id":"UCitFTWVsXPNlsXHFtR8v24A",
            "title":"OneHour",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLSWYhfjnu9dly3tan77xUMlligP3DNS1pfNbi09=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCitFTWVsXPNlsXHFtR8v24A"
          },
          "url":"https://www.youtube.com/watch?v=6vztdGgJaAw"
        },
        {
          "id":"SEsvsSaJadQ",
          "title":"Janji - Heroes Tonight (10 Hours)",
          "publishedTime":"3 years ago",
          "duration":36004,
          "viewCount":2770382,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/SEsvsSaJadQ/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLBHsiodwfiaFFBmxiHwSi6N4xLwjg",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/SEsvsSaJadQ/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLA5u6CyMa8aX8Z_fRcVdfEh8kQnNw",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/SEsvsSaJadQ/mqdefault_6s.webp?du=3000\u0026sqp=CLjJuIkG\u0026rs=AOn4CLDFf3T5BEncUGeKEpfIaWNJWPP6Ew",
            "height":180,
            "width":320
          },
          "description":"Would you like to support me? https://www.patreon.com/10Hours Janji SoundCloud https://soundcloud.com/janjimusic Facebook ...",
          "channel":{
            "id":"UCOJuA_374WrqK53ysxstNBQ",
            "title":"Totej",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLSMWMIRRoCUBN1jlYpsai7cuFbjXtJSDjl19vRMPw=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCOJuA_374WrqK53ysxstNBQ"
          },
          "url":"https://www.youtube.com/watch?v=SEsvsSaJadQ"
        },
        {
          "id":"kJ98U74xn5Q",
          "title":"NCS: 30 Million Subscriber Mix | Remake",
          "publishedTime":"3 months ago",
          "duration":5727,
          "viewCount":202856,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/kJ98U74xn5Q/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLDpxBI1VQVxH7MQp3gLK4k_rHsNkg",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/kJ98U74xn5Q/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCpTwir-V5KMkN17p0lfI6IFNN-5Q",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/kJ98U74xn5Q/mqdefault_6s.webp?du=3000\u0026sqp=CJDmuIkG\u0026rs=AOn4CLDJL5BGPaLdCfKrR47vx2wBXsCElg",
            "height":180,
            "width":320
          },
          "description":"bet u didn't see this coming! :) Original video: https://youtu.be/JNl1_hRwpXE Track list: 00:00​ Cartoon - On \u0026 On (feat.",
          "channel":{
            "id":"UCYWRKABK63tI5s_RXR9xhCg",
            "title":"Foxster",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLSX7AfWEUY9WRFr0MIan_xXNBK-rwbf0OQM_y-GBQ=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCYWRKABK63tI5s_RXR9xhCg"
          },
          "url":"https://www.youtube.com/watch?v=kJ98U74xn5Q"
        },
        {
          "id":"pLZq3jgE6qA",
          "title":"Top 20 songs of Tobu - Best Of Tobu",
          "publishedTime":"5 years ago",
          "duration":4665,
          "viewCount":12083805,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/pLZq3jgE6qA/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCbGXE8dk4WZpOSPucGv4umF_te6A",
              "height":202,
              "width":360
            },
            {
              "url":"https://i.ytimg.com/vi/pLZq3jgE6qA/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLC-mA-laVOH00NQBi7uIW-remiYyw",
              "height":404,
              "width":720
            }
          ],
          "richThumbnail":{
            "url":"https://i.ytimg.com/an_webp/pLZq3jgE6qA/mqdefault_6s.webp?du=3000\u0026sqp=CIDYuIkG\u0026rs=AOn4CLCBtvF4po4w0O5InNKFajS58bawxw",
            "height":180,
            "width":320
          },
          "description":"Top 20 songs of Tobu - Best Of Tobu Tobu Collection ☞ Follow Magic NCS : Facebook → https://goo.gl/fWA96C Subscribe ...",
          "channel":{
            "id":"UCoDZIZuadPBixSPFR7jAq2A",
            "title":"Music Store",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLSNH-mOdxWiiqLGpFJAL6Zq87MvN7w9I42Hr9a_=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCoDZIZuadPBixSPFR7jAq2A"
          },
          "url":"https://www.youtube.com/watch?v=pLZq3jgE6qA"
        },
        {
          "id":"nzXIzydwDxs",
          "title":"Ncs full album terbaik",
          "publishedTime":"9 months ago",
          "duration":2520,
          "viewCount":291270,
          "thumbnails":[
            {
              "url":"https://i.ytimg.com/vi/nzXIzydwDxs/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCQQXPrPtRPen1roAF6VnA-y5L5qQ",
              "height":270,
              "width":480
            }
          ],
          "richThumbnail":{
            "url":"",
            "height":0,
            "width":0
          },
          "description":"ncs #nocopyrightsounds.",
          "channel":{
            "id":"UCcK1C246TG_OQbn2dsZ4vpQ",
            "title":"Lagu Ku",
            "thumbnails":[
              {
                "url":"https://yt3.ggpht.com/ytc/AKedOLR_zdp_16YFfvUJdiYtQ5y5R3RPRxAL6WSz4Q=s68-c-k-c0x00ffffff-no-rj",
                "height":68,
                "width":68
              }
            ],
            "url":"https://www.youtube.com/channel/UCcK1C246TG_OQbn2dsZ4vpQ"
          },
          "url":"https://www.youtube.com/watch?v=nzXIzydwDxs"
        }
      ]
    }
  ],
  "suggestions":[
    "ncs spectre",
    "ncs 1 hour",
    "ncs tobu",
    "ncs background music",
    "ncs heroes tonight",
    "ncs mortals",
    "ncs fearless",
    "ncs invincible",
    "ncs superhero",
    "ncs songs for gaming",
    "ncs fade",
    "ncs sky high",
    "ncs instrumental",
    "ncs alan walker",
    "ncs hindi songs",
    "ncs vlog music",
    "ncs music bass boosted",
    "ncs ringtone"
  ]
}
```
</details>

### Use filters and sorts

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/raitonoberu/ytsearch"
)

func main() {
    search := ytsearch.Search("ncs")

    search.SearchFilter = ytsearch.VideoFilter
    search.SortOrder = ytsearch.ViewCountOrder

    result, err := search.Next()
    if err != nil {
        panic(err)
    }

    jsonstr, _ := json.Marshal(result.Videos[0])
    fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
  "id":"bM7SZ5SBzyY",
  "title":"Alan Walker - Fade [NCS Release]",
  "publishedTime":"6 years ago",
  "duration":261,
  "viewCount":444680839,
  "thumbnails":[
    {
      "url":"https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==\u0026rs=AOn4CLCndJeieOK_hJwZkKKLRI4mm7uHtA",
      "height":270,
      "width":480
    }
  ],
  "richThumbnail":{
    "url":"https://i.ytimg.com/an_webp/bM7SZ5SBzyY/mqdefault_6s.webp?du=3000\u0026sqp=CPTLuIkG\u0026rs=AOn4CLBTZTjoBSbEHeKlNScelYyCSy0WWg",
    "height":180,
    "width":320
  },
  "description":"NCS: Music Without Limitations NCS Spotify: http://spoti.fi/NCS Alan Walker - Fade is included in our debut compilation NCS: ...",
  "channel":{
    "id":"UC_aEa8K-EOJ3D6gOs7HcyNg",
    "title":"NoCopyrightSounds",
    "thumbnails":[
      {
        "url":"https://yt3.ggpht.com/ytc/AKedOLQPUsOFTh7JnkDWMwnS737uhnk7EsENZkBXnZ5rMA=s68-c-k-c0x00ffffff-no-rj",
        "height":68,
        "width":68
      }
    ],
    "url":"https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
  },
  "url":"https://www.youtube.com/watch?v=bM7SZ5SBzyY"
}
```
</details>


## License

**MIT License**, see [LICENSE](./LICENSE) file for additional information.
