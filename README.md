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
	search.Limit = 2
	result, err := search.Next()
	if err != nil {
		panic(err)
	}

	jsonstr, _ := json.Marshal(result.Videos)
	fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
[
  {
    "id": "-ObdvMkCKws",
    "title": "Top 20 Most Popular Songs by NCS | Best of NCS | Most Viewed Songs",
    "published_time": "1 year ago",
    "duration": 4367,
    "view_count": 25793257,
    "thumbnails": [
      {
        "url": "https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCXso3cln-2aZPAiSNW15njQ-xchw",
        "height": 202,
        "width": 360
      },
      {
        "url": "https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCYGQ4cDZ49e0mlnD2Uj_oAI03wzQ",
        "height": 404,
        "width": 720
      }
    ],
    "rich_thumbnail": null,
    "description": "Top 20 Most Popular Songs by NCS | Best of NCS | Most Viewed Songs Tracklist: 0:00:00 – Janji – Heroes Tonight (feat.",
    "channel": {
      "id": "UCoDZIZuadPBixSPFR7jAq2A",
      "title": "Music Store",
      "thumbnails": [
        {
          "url": "https://yt3.ggpht.com/ytc/AAUvwng5Qb-WbINYd4rajHdqppVduiHmG8h1q8UhVeC0=s68-c-k-c0x00ffffff-no-rj",
          "height": 68,
          "width": 68
        }
      ],
      "link": "https://www.youtube.com/channel/UCoDZIZuadPBixSPFR7jAq2A"
    },
    "link": "https://www.youtube.com/watch?v=-ObdvMkCKws"
  },
  {
    "id": "mIrbvjZRE9c",
    "title": "Slippy & Blosso - Horizon (Back To Life) (Feat. GLNNA) [NCS Release]",
    "published_time": "1 day ago",
    "duration": 200,
    "view_count": 265863,
    "thumbnails": [
      {
        "url": "https://i.ytimg.com/vi/mIrbvjZRE9c/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDpESN04BeQ6oROxQLw_PlvTCuT4A",
        "height": 202,
        "width": 360
      },
      {
        "url": "https://i.ytimg.com/vi/mIrbvjZRE9c/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLARDBKik976qluuvL-Aoh6HZh5vrg",
        "height": 404,
        "width": 720
      }
    ],
    "rich_thumbnail": {
      "url": "https://i.ytimg.com/an_webp/mIrbvjZRE9c/mqdefault_6s.webp?du=3000&sqp=COWqnIYG&rs=AOn4CLBDozoF4xuO_jtDOr661ZbMQ64ojw",
      "height": 180,
      "width": 320
    },
    "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
    "channel": {
      "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
      "title": "NoCopyrightSounds",
      "thumbnails": [
        {
          "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
          "height": 68,
          "width": 68
        }
      ],
      "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
    },
    "link": "https://www.youtube.com/watch?v=mIrbvjZRE9c"
  }
]
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
	search.Limit = 2
	result, err := search.Next()
	if err != nil {
		panic(err)
	}

	jsonstr, _ := json.Marshal(result.Channels)
	fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
[
  {
    "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
    "title": "NoCopyrightSounds",
    "thumbnails": [
      {
        "url": "http://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s88-c-k-c0x00ffffff-no-rj-mo",
        "height": 88,
        "width": 88
      },
      {
        "url": "http://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s176-c-k-c0x00ffffff-no-rj-mo",
        "height": 176,
        "width": 176
      }
    ],
    "videoCount": 931,
    "description": "NoCopyrightSounds is a copyright free / stream safe record label, providing free to use music to the content creator community.",
    "subscribers": "30.5M subscribers",
    "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
  },
  {
    "id": "UC7P9rGN3iFD_LLK9eI2kqAg",
    "title": "NCS Music/ Музыка без авторских прав",
    "thumbnails": [
      {
        "url": "http://yt3.ggpht.com/ytc/AAUvwnhdM7Rgz8lIcp3BkH00pMDNf8T_ntnfjcTmCV71=s88-c-k-c0x00ffffff-no-rj-mo",
        "height": 88,
        "width": 88
      },
      {
        "url": "http://yt3.ggpht.com/ytc/AAUvwnhdM7Rgz8lIcp3BkH00pMDNf8T_ntnfjcTmCV71=s176-c-k-c0x00ffffff-no-rj-mo",
        "height": 176,
        "width": 176
      }
    ],
    "videoCount": 143,
    "description": "На канале вы можете найти самые новые и популярные треки без авторских прав, для работы с Youtube! Если вы ...",
    "subscribers": "24.5K subscribers",
    "link": "https://www.youtube.com/channel/UC7P9rGN3iFD_LLK9eI2kqAg"
  }
]
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
	search.Limit = 2
	result, err := search.Next()
	if err != nil {
		panic(err)
	}

	jsonstr, _ := json.Marshal(result.Playlists)
	fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
[
  {
    "id": "PLUujrJZl_60rr9OQMLSzHvrbIX0dCen-i",
    "title": "NCS Classics",
    "videoCount": 31,
    "channel": {
      "id": "UCO5mPtqSHtb-yNtDnZsV_iQ",
      "title": "Hipstersky Drwal",
      "link": "https://www.youtube.com/channel/UCO5mPtqSHtb-yNtDnZsV_iQ"
    },
    "thumbnails": [
      {
        "url": "http:https://i.ytimg.com/vi/p7ZsBPK656s/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLApeIroWtn-fdFVS7d7-opZUokw4g",
        "height": 94,
        "width": 168
      },
      {
        "url": "http:https://i.ytimg.com/vi/p7ZsBPK656s/hqdefault.jpg?sqp=-oaymwEWCMQBEG5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLDHIM9tWQgbTw14LaURg8G8Bs0xCw",
        "height": 110,
        "width": 196
      },
      {
        "url": "http:https://i.ytimg.com/vi/p7ZsBPK656s/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLBijJd3rKt18bDDtyOUGVe1yRfn-g",
        "height": 138,
        "width": 246
      },
      {
        "url": "http:https://i.ytimg.com/vi/p7ZsBPK656s/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLBXRKtFd5dFGmBETs_IOHLvy4GTcg",
        "height": 188,
        "width": 336
      }
    ],
    "link": "https://www.youtube.com/playlist?list=PLUujrJZl_60rr9OQMLSzHvrbIX0dCen-i"
  },
  {
    "id": "PLhy8TB5U6n17R78U7usaLQfCC8nbnG8Nc",
    "title": "Best of NCS Playlist",
    "videoCount": 195,
    "channel": {
      "id": "UCtMN6g3-VLVHegoAFH3y_Sg",
      "title": "오진서",
      "link": "https://www.youtube.com/channel/UCtMN6g3-VLVHegoAFH3y_Sg"
    },
    "thumbnails": [
      {
        "url": "http:https://i.ytimg.com/vi/iqoNoU-rm14/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLBML48rBgPw04Se6JrEBjgJlUn7-g",
        "height": 94,
        "width": 168
      },
      {
        "url": "http:https://i.ytimg.com/vi/iqoNoU-rm14/hqdefault.jpg?sqp=-oaymwEWCMQBEG5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLAoErTSbsty1nhSl1mLgKqE-mjbXQ",
        "height": 110,
        "width": 196
      },
      {
        "url": "http:https://i.ytimg.com/vi/iqoNoU-rm14/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLBHMrUe74r4QaEhWKHkrI2rmlK19A",
        "height": 138,
        "width": 246
      },
      {
        "url": "http:https://i.ytimg.com/vi/iqoNoU-rm14/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCu37G726mlZij3h0ldmA8QXTiVKg",
        "height": 188,
        "width": 336
      }
    ],
    "link": "https://www.youtube.com/playlist?list=PLhy8TB5U6n17R78U7usaLQfCC8nbnG8Nc"
  }
]
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
  "estimated_results": 27946758,
  "videos": [
    {
      "id": "-ObdvMkCKws",
      "title": "Top 20 Most Popular Songs by NCS | Best of NCS | Most Viewed Songs",
      "published_time": "1 year ago",
      "duration": 4367,
      "view_count": 25793257,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCXso3cln-2aZPAiSNW15njQ-xchw",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/-ObdvMkCKws/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCYGQ4cDZ49e0mlnD2Uj_oAI03wzQ",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": null,
      "description": "Top 20 Most Popular Songs by NCS | Best of NCS | Most Viewed Songs Tracklist: 0:00:00 – Janji – Heroes Tonight (feat.",
      "channel": {
        "id": "UCoDZIZuadPBixSPFR7jAq2A",
        "title": "Music Store",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwng5Qb-WbINYd4rajHdqppVduiHmG8h1q8UhVeC0=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UCoDZIZuadPBixSPFR7jAq2A"
      },
      "link": "https://www.youtube.com/watch?v=-ObdvMkCKws"
    },
    {
      "id": "eUR-DrTGNws",
      "title": "Top 10 Most Popular Songs by NCS | Episode 1",
      "published_time": "7 months ago",
      "duration": 2404,
      "view_count": 12247380,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/eUR-DrTGNws/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD0A-CGEUwsgAWfryyOArKag_GUXw",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/eUR-DrTGNws/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCqaT6eE5C0eNJNgDkCD9vrE7p5IQ",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/eUR-DrTGNws/mqdefault_6s.webp?du=3000&sqp=CJ6snIYG&rs=AOn4CLC31JO0fsTnHMPk0GzCk7JAbNvi7w",
        "height": 180,
        "width": 320
      },
      "description": "0:00 - 1. Alan Walker - Fade 4:20 - 2. Cartoon - On & On 7:48 - 3. Alan Walker - Spectre 11:34 - 4. DEAF KEV - Invincible 16:07 - 5.",
      "channel": {
        "id": "UC0raGe2owvmadTf85xL_ytw",
        "title": "Limon NCS",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwniGkPQtxNuubnnVhIj0eVcdqkZ20EXUnSXBP6VH=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC0raGe2owvmadTf85xL_ytw"
      },
      "link": "https://www.youtube.com/watch?v=eUR-DrTGNws"
    },
    {
      "id": "JNl1_hRwpXE",
      "title": "NCS: 30 Million Subscriber Mix",
      "published_time": "1 month ago",
      "duration": 5726,
      "view_count": 2827866,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/JNl1_hRwpXE/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDwkR7V56OMD5UY9uwyr1ImIAu97Q",
          "height": 270,
          "width": 480
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/JNl1_hRwpXE/mqdefault_6s.webp?du=3000&sqp=CJiRnIYG&rs=AOn4CLBeKah-7P6YUd49CsN_HnUaVFeRyQ",
        "height": 180,
        "width": 320
      },
      "description": "Best of NoCopyrightSounds 2021 (30 Million Subscriber Mix) WE DID IT!! ❤️ Thank you SO much for all the support you guys ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=JNl1_hRwpXE"
    },
    {
      "id": "pu31wuUTfRI",
      "title": "🔥Artistic Music Mix: Top 30 Songs ♫ Best NCS Gaming Music 2021 Mix ♫ EDM, Trap, DnB, Dubstep, House",
      "published_time": "3 hours ago",
      "duration": 10800,
      "view_count": 628,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/pu31wuUTfRI/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD_RHzfxT9OnwR3w5V2xHNq3NjoMg",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/pu31wuUTfRI/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCOX9qDzUBlWLj5xQRYiqK6GcF8iw",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/pu31wuUTfRI/mqdefault_6s.webp?du=3000&sqp=COSynIYG&rs=AOn4CLAljoGYlUxBs_ZBvTf1Gog7cFeHuw",
        "height": 180,
        "width": 320
      },
      "description": "Welcome to Freeme NCS☆ ------- ▻Artistic Music Mix: Top 30 Songs ♫ Best NCS Gaming Music 2021 Mix ♫ EDM, Trap, DnB, ...",
      "channel": {
        "id": "UCG0SzK_t4-Ylf1yZq9Xmi_g",
        "title": "Freeme NCS Music",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwngkwQtsY2wvJucAQ2r3HQHlYfLkGXN5ulPHyNUB=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UCG0SzK_t4-Ylf1yZq9Xmi_g"
      },
      "link": "https://www.youtube.com/watch?v=pu31wuUTfRI"
    },
    {
      "id": "yJg-Y5byMMw",
      "title": "Warriyo - Mortals (feat. Laura Brehm) [NCS Release]",
      "published_time": "4 years ago",
      "duration": 230,
      "view_count": 180173254,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/yJg-Y5byMMw/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAmKIY6QpyYwke2y20r8AbLNnEQjw",
          "height": 270,
          "width": 480
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/yJg-Y5byMMw/mqdefault_6s.webp?du=3000&sqp=CJC-nIYG&rs=AOn4CLA_JhkG9YRekbwIvkGkVGOlkxS5Vw",
        "height": 180,
        "width": 320
      },
      "description": "NCS: Music Without Limitations NCS Spotify: http://spoti.fi/NCS Free Download / Stream: http://ncs.io/mortals Connect with NCS: ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=yJg-Y5byMMw"
    },
    {
      "id": "ABuNwLP-z9o",
      "title": "🔥 Top 50 NoCopyRightSounds | Best of NCS | Most viewed ! Gaming Music | The Best of All Time | 2021",
      "published_time": "2 years ago",
      "duration": 11026,
      "view_count": 9535569,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/ABuNwLP-z9o/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLC0rl3DGDwIvQs3ADcVp40EyQmhDw",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/ABuNwLP-z9o/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBUWW_Gz6kYwRcNJ9yudEipM2idwQ",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/ABuNwLP-z9o/mqdefault_6s.webp?du=3000&sqp=COqfnIYG&rs=AOn4CLDGlee12gaft6HTKVKPp4SjvZXzhg",
        "height": 180,
        "width": 320
      },
      "description": "Best of NCS | The Best of All Time Track list: 01. 0:00 Alan Walker - Fade 02. 4:20 Alan Walker - Spectre 03. 8:07 Different Heaven ...",
      "channel": {
        "id": "UCG0SzK_t4-Ylf1yZq9Xmi_g",
        "title": "Freeme NCS Music",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwngkwQtsY2wvJucAQ2r3HQHlYfLkGXN5ulPHyNUB=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UCG0SzK_t4-Ylf1yZq9Xmi_g"
      },
      "link": "https://www.youtube.com/watch?v=ABuNwLP-z9o"
    },
    {
      "id": "S19UcWdOA-I",
      "title": "Lost Sky - Fearless pt.II (feat. Chris Linton) [NCS Release]",
      "published_time": "3 years ago",
      "duration": 194,
      "view_count": 123340869,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/S19UcWdOA-I/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDMiju4cvt7r4RId6P3lLQYUepbpg",
          "height": 270,
          "width": 480
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/S19UcWdOA-I/mqdefault_6s.webp?du=3000&sqp=CPKlnIYG&rs=AOn4CLBVGOG3ZbJcX-zcXaZMRFPWtYNfJw",
        "height": 180,
        "width": 320
      },
      "description": "- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=S19UcWdOA-I"
    },
    {
      "id": "8AyHjrQYszU",
      "title": "💥Awesome Gaming Mix: Top 30 Songs ♫ Best NCS Gaming Music 2021 Mix ♫ EDM, Trap, DnB, Dubstep, House",
      "published_time": "2 days ago",
      "duration": 10800,
      "view_count": 34824,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/8AyHjrQYszU/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLC8lOKhiI09AuGivPC5MCIM8LDwpA",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/8AyHjrQYszU/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCHxMAAyohA1wZ4CQ8KrVg66tInkA",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/8AyHjrQYszU/mqdefault_6s.webp?du=3000&sqp=CPKlnIYG&rs=AOn4CLA4V9zPXEGfPVbW_-uyPOQRBesTRQ",
        "height": 180,
        "width": 320
      },
      "description": "Welcome to Freeme NCS☆ ------- ▻Awesome Gaming Mix: Top 30 Songs ♫ Best NCS Gaming Music 2021 Mix ♫ EDM, Trap, ...",
      "channel": {
        "id": "UCG0SzK_t4-Ylf1yZq9Xmi_g",
        "title": "Freeme NCS Music",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwngkwQtsY2wvJucAQ2r3HQHlYfLkGXN5ulPHyNUB=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UCG0SzK_t4-Ylf1yZq9Xmi_g"
      },
      "link": "https://www.youtube.com/watch?v=8AyHjrQYszU"
    },
    {
      "id": "mj9KRKSvdbk",
      "title": "Clarx - H.A.Y [NCS Release]",
      "published_time": "3 years ago",
      "duration": 224,
      "view_count": 33263291,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/mj9KRKSvdbk/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLC84-_ahhkbpBmDGx1YjRMpcMqkeA",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/mj9KRKSvdbk/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDCWgiiuzjgQ_eC01VtS6CwAvvKZw",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/mj9KRKSvdbk/mqdefault_6s.webp?du=3000&sqp=CLWnnIYG&rs=AOn4CLBggpGR0j_u1aL8DlkScNrm_RuNMA",
        "height": 180,
        "width": 320
      },
      "description": "- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=mj9KRKSvdbk"
    },
    {
      "id": "MEkaqZecpUQ",
      "title": "NCS: 2019 ‘20 Million’ Mix | Future Hits",
      "published_time": "2 years ago",
      "duration": 3634,
      "view_count": 14553802,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/MEkaqZecpUQ/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD7jTWOl0sO0KveSWI_1bstCx762w",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/MEkaqZecpUQ/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAIoLSCY2DkOkejEB-Wm9sLi7Llrg",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/MEkaqZecpUQ/mqdefault_6s.webp?du=3000&sqp=COiznIYG&rs=AOn4CLDuiUPUVh37RKJv0r_puw67Yt0LEA",
        "height": 180,
        "width": 320
      },
      "description": "2019 Best of NCS '20 Million Subscriber' Mix! Thank you all for 20 MILLION subscribers! Here is a mix of some classics we've ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=MEkaqZecpUQ"
    },
    {
      "id": "J2X5mJ3HDYE",
      "title": "DEAF KEV - Invincible [NCS Release]",
      "published_time": "6 years ago",
      "duration": 274,
      "view_count": 245811260,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/J2X5mJ3HDYE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAt3Ibnfb7NuT7S9b7wD8oIlhZbTQ",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/J2X5mJ3HDYE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLC_LiPd8jGeALvV1sFhbNxt818qxQ",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/J2X5mJ3HDYE/mqdefault_6s.webp?du=3000&sqp=CIasnIYG&rs=AOn4CLAm3Ful7Vsw5RifGEprUQMqTD2Nlw",
        "height": 180,
        "width": 320
      },
      "description": "NCS: Music Without Limitations NCS Spotify: http://spoti.fi/NCS Free Download / Stream: http://ncs.io/invincible ▽ Connect with ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=J2X5mJ3HDYE"
    },
    {
      "id": "K4DyBUG242c",
      "title": "Cartoon - On & On (feat. Daniel Levi) [NCS Release]",
      "published_time": "5 years ago",
      "duration": 208,
      "view_count": 418877980,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/K4DyBUG242c/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCVDGZSQEPvPuF1QyELmT2FV6vdaQ",
          "height": 270,
          "width": 480
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/K4DyBUG242c/mqdefault_6s.webp?du=3000&sqp=CLSxnIYG&rs=AOn4CLD09OYGOdDGp81ZNimc0K3cGVX_hA",
        "height": 180,
        "width": 320
      },
      "description": "Lyrics: Cartoon - On & On feat. Daniel Levi [Verse 1] Hold me close 'til I get up Time is barely on our side I don't wanna waste ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=K4DyBUG242c"
    },
    {
      "id": "C5fLxtJH2Qs",
      "title": "Egzod & Maestro Chives - Royalty (ft. Neoni) [NCS Release]",
      "published_time": "1 month ago",
      "duration": 224,
      "view_count": 1934769,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/C5fLxtJH2Qs/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDY1smnn9PHE2eLHAuWFtVlKZ3vvQ",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/C5fLxtJH2Qs/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAxwLHjOgAxes1042xYKq8bl_oV1A",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/C5fLxtJH2Qs/mqdefault_6s.webp?du=3000&sqp=CIeZnIYG&rs=AOn4CLBaJ6UscLx00bmBk0m704KG7rXG9g",
        "height": 180,
        "width": 320
      },
      "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=C5fLxtJH2Qs"
    },
    {
      "id": "49aqYW1GVZU",
      "title": "Топ 20 самых популярных песен | Лучший из NCS | Самые популярные песни",
      "published_time": "2 years ago",
      "duration": 4446,
      "view_count": 96146,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/49aqYW1GVZU/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAvvbpZpPDBErfjGhwCV3u66t19eA",
          "height": 270,
          "width": 480
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/49aqYW1GVZU/mqdefault_6s.webp?du=3000&sqp=CJG4nIYG&rs=AOn4CLC3zpBsIbeYVz8UaUF2UYufOSsdxQ",
        "height": 180,
        "width": 320
      },
      "description": "Топ 20 самых популярных песен NCS | Лучший из NCS | Самые популярные песни ♫ ◅ Это музыкальный микс NCS, ...",
      "channel": {
        "id": "UCUfNUW5j_TnnHnpPYq-KuhQ",
        "title": "insane TOP",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnh76iOJwgUlqVUAPcUkHQLEQcQs9jgcj9L7z-x9=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UCUfNUW5j_TnnHnpPYq-KuhQ"
      },
      "link": "https://www.youtube.com/watch?v=49aqYW1GVZU"
    },
    {
      "id": "L7kF4MXXCoA",
      "title": "Lost Sky - Dreams pt. II (feat. Sara Skinner) [NCS Release]",
      "published_time": "2 years ago",
      "duration": 215,
      "view_count": 59101638,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/L7kF4MXXCoA/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBSpvILy5tmybx9SbeipbEPnraMvQ",
          "height": 270,
          "width": 480
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/L7kF4MXXCoA/mqdefault_6s.webp?du=3000&sqp=CLC6nIYG&rs=AOn4CLDAiFnA3hflgtVn13Ykuh-aO2wiEg",
        "height": 180,
        "width": 320
      },
      "description": "- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - When you are using this track, we simply ask that you put this in your description: ...",
      "channel": {
        "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
        "title": "NoCopyrightSounds",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
      },
      "link": "https://www.youtube.com/watch?v=L7kF4MXXCoA"
    },
    {
      "id": "7tNtU5XFwrU",
      "title": "NCS 24/7 - Copyright Free Music Livestream by NoCopyrightSounds",
      "published_time": "",
      "duration": 0,
      "view_count": 0,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/7tNtU5XFwrU/hq720_live.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCixG8NhGZcK2Jrno9FQhKBx4FtPw",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/7tNtU5XFwrU/hq720_live.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLA0-Xe_tzc1EOsk6s4rRffdMnTDQw",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": null,
      "description": "The best in copyright-free music from NoCopyrightSounds while you're gaming, studying or just relaxing... we've got you covered!",
      "channel": {
        "id": "UCiJnBO_XuDsi1SSRAmt4n5g",
        "title": "NCS Arcade",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwnjzirgQDkB6tQkL5uApSnPGMEhR_xEuapBh-xdU=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UCiJnBO_XuDsi1SSRAmt4n5g"
      },
      "link": "https://www.youtube.com/watch?v=7tNtU5XFwrU"
    },
    {
      "id": "PcqeF3nNvWo",
      "title": "⚡Inspiring Music: Top 30 Vocal Mix ♫ Best NCS Gaming Music 2021 Mix ♫ EDM, Trap, DnB, Dubstep, House",
      "published_time": "3 days ago",
      "duration": 10800,
      "view_count": 15262,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/PcqeF3nNvWo/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLA8G2cZ8ecvxnlqIRUVzexGUIHyag",
          "height": 202,
          "width": 360
        },
        {
          "url": "https://i.ytimg.com/vi/PcqeF3nNvWo/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAp5x9SswBEKhOvqrDoKV-MHRDMjw",
          "height": 404,
          "width": 720
        }
      ],
      "rich_thumbnail": {
        "url": "https://i.ytimg.com/an_webp/PcqeF3nNvWo/mqdefault_6s.webp?du=3000&sqp=CLaNnIYG&rs=AOn4CLDgnXJDqUOmv2A8FbyLLA0K5zYy7g",
        "height": 180,
        "width": 320
      },
      "description": "Welcome to Freeme NCS☆ ------- ▻Inspiring Music: Top 30 Vocal Mix ♫ Best NCS Gaming Music 2021 Mix ♫ EDM, Trap, DnB, ...",
      "channel": {
        "id": "UCG0SzK_t4-Ylf1yZq9Xmi_g",
        "title": "Freeme NCS Music",
        "thumbnails": [
          {
            "url": "https://yt3.ggpht.com/ytc/AAUvwngkwQtsY2wvJucAQ2r3HQHlYfLkGXN5ulPHyNUB=s68-c-k-c0x00ffffff-no-rj",
            "height": 68,
            "width": 68
          }
        ],
        "link": "https://www.youtube.com/channel/UCG0SzK_t4-Ylf1yZq9Xmi_g"
      },
      "link": "https://www.youtube.com/watch?v=PcqeF3nNvWo"
    }
  ],
  "channels": [
    {
      "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
      "title": "NoCopyrightSounds",
      "thumbnails": [
        {
          "url": "http://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s88-c-k-c0x00ffffff-no-rj-mo",
          "height": 88,
          "width": 88
        },
        {
          "url": "http://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s176-c-k-c0x00ffffff-no-rj-mo",
          "height": 176,
          "width": 176
        }
      ],
      "videoCount": 931,
      "description": "NoCopyrightSounds is a copyright free / stream safe record label, providing free to use music to the content creator community.",
      "subscribers": "30.5M subscribers",
      "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
    }
  ],
  "playlists": [
    {
      "id": "PLzkuLC6Yvumv_Rd5apfPRWEcjf9b1JRnq",
      "title": "NCS Playlist Best Songs",
      "videoCount": 792,
      "channel": {
        "id": "UC7o5WwsbRAdfZwDe2brud7g",
        "title": "NCS",
        "link": "https://www.youtube.com/channel/UC7o5WwsbRAdfZwDe2brud7g"
      },
      "thumbnails": [
        {
          "url": "http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLAD6Vv5Z0roTmEwnNF_NyBhfJKN-w",
          "height": 94,
          "width": 168
        },
        {
          "url": "http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEWCMQBEG5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLDxRqXhHvPhehwYkIkGhQkGx9OswQ",
          "height": 110,
          "width": 196
        },
        {
          "url": "http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCcakX_wGhCA1InSGO01CTTVn0AdQ",
          "height": 138,
          "width": 246
        },
        {
          "url": "http:https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCcWVWWwwNAxJ_F5NOVwGXepJxsQA",
          "height": 188,
          "width": 336
        }
      ],
      "link": "https://www.youtube.com/playlist?list=PLzkuLC6Yvumv_Rd5apfPRWEcjf9b1JRnq"
    }
  ],
  "shelves": [
    {
      "title": "Latest from NoCopyrightSounds",
      "items": [
        {
          "id": "mIrbvjZRE9c",
          "title": "Slippy & Blosso - Horizon (Back To Life) (Feat. GLNNA) [NCS Release]",
          "published_time": "1 day ago",
          "duration": 200,
          "view_count": 270240,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/mIrbvjZRE9c/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDpESN04BeQ6oROxQLw_PlvTCuT4A",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/mIrbvjZRE9c/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLARDBKik976qluuvL-Aoh6HZh5vrg",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/mIrbvjZRE9c/mqdefault_6s.webp?du=3000&sqp=COWqnIYG&rs=AOn4CLBDozoF4xuO_jtDOr661ZbMQ64ojw",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=mIrbvjZRE9c"
        },
        {
          "id": "ZbnMeTzzjZE",
          "title": "Alisky - Grow (feat. VØR) [NCS Release]",
          "published_time": "3 days ago",
          "duration": 153,
          "view_count": 382619,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/ZbnMeTzzjZE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCvwp00NUgOSmsKnT11UsHVEHlRnQ",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/ZbnMeTzzjZE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCIagsDbS98xdVgKQJ-lVegcp0cww",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/ZbnMeTzzjZE/mqdefault_6s.webp?du=3000&sqp=CPCrnIYG&rs=AOn4CLAX-noF5eLOMZvn45-xM1QETQm-xg",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=ZbnMeTzzjZE"
        },
        {
          "id": "QoxnZ0tAH_A",
          "title": "Rameses B - Hardwired [NCS Release]",
          "published_time": "5 days ago",
          "duration": 302,
          "view_count": 482223,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/QoxnZ0tAH_A/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLA3wzKE0v-OH4wGkhTX17uSI1sTNA",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/QoxnZ0tAH_A/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBZTa5iwoFIGRs228frV70IwfZJOg",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/QoxnZ0tAH_A/mqdefault_6s.webp?du=3000&sqp=CJCInIYG&rs=AOn4CLDnhmep7ZcYofjpCIV8FcDlwf61aA",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=QoxnZ0tAH_A"
        },
        {
          "id": "nulzVedhSRE",
          "title": "Facading - Freefalling  [NCS Release]",
          "published_time": "1 week ago",
          "duration": 154,
          "view_count": 599224,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/nulzVedhSRE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAybxuF9kJG_AF2Is4EUoTE2pFqcA",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/nulzVedhSRE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDePxvwciqNf1lsIknN2fd1I2t4tw",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/nulzVedhSRE/mqdefault_6s.webp?du=3000&sqp=CIu8nIYG&rs=AOn4CLDmoLt6Hg76GxtLd7-vqF-e-ViU9Q",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=nulzVedhSRE"
        },
        {
          "id": "P__8oculr0o",
          "title": "Abandoned & GalaxyTones - Luna (Feat. DNAKM) [NCS Release]",
          "published_time": "1 week ago",
          "duration": 201,
          "view_count": 502494,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/P__8oculr0o/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCd3E1-7nmUibMJhwMg8Ojs-rtO1w",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/P__8oculr0o/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDZnFUYMbAqJ90aOIx8LEZe74s2bg",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/P__8oculr0o/mqdefault_6s.webp?du=3000&sqp=CIqxnIYG&rs=AOn4CLAhQ8B567GzwBde9vNEGOwqkGFiig",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=P__8oculr0o"
        },
        {
          "id": "KVbrHTRxgsk",
          "title": "More Plastic - Rewind  [NCS Release]",
          "published_time": "1 week ago",
          "duration": 204,
          "view_count": 575249,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/KVbrHTRxgsk/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLA-k25KoVAwbL79eJ8ttw7092Z3JA",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/KVbrHTRxgsk/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDTHwJfHX8snNAjw5MrmoP1LXyyTA",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/KVbrHTRxgsk/mqdefault_6s.webp?du=3000&sqp=CN63nIYG&rs=AOn4CLAkczhwxv4pvEV2x3WeXC-ajtRsUw",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=KVbrHTRxgsk"
        },
        {
          "id": "O9UgaCWN2Ag",
          "title": "Clarx & Moe Aly - Healing [NCS Release]",
          "published_time": "2 weeks ago",
          "duration": 182,
          "view_count": 724981,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/O9UgaCWN2Ag/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBq-EUMYsNVamJ4gPrZpeeVAI2W1Q",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/O9UgaCWN2Ag/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCJtq0yN7k7oxhqV_ZIl5S2ktPt5w",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/O9UgaCWN2Ag/mqdefault_6s.webp?du=3000&sqp=CICcnIYG&rs=AOn4CLA7xQ_xk02dESXqqrBDYvrc8LjFdQ",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=O9UgaCWN2Ag"
        },
        {
          "id": "woNrNfFVsKc",
          "title": "CHENDA & Shiah Maisel - Ten More Minutes  [NCS Release]",
          "published_time": "2 weeks ago",
          "duration": 222,
          "view_count": 557972,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/woNrNfFVsKc/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCO1QwYoqeH8Fj3DwH3VXlJ9p6xow",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/woNrNfFVsKc/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBJclgvuvOVmscnw6x55Yy3VQiNyw",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/woNrNfFVsKc/mqdefault_6s.webp?du=3000&sqp=CKivnIYG&rs=AOn4CLD89fbhxjXDsva95jaMrYkunQtn2g",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=woNrNfFVsKc"
        },
        {
          "id": "I9_xOTF3DiE",
          "title": "if found - Need You [NCS Release]",
          "published_time": "2 weeks ago",
          "duration": 197,
          "view_count": 704707,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/I9_xOTF3DiE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBqqK3aOzVsVKcX8ybGeBalLwo9Bw",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/I9_xOTF3DiE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDlI6rW3oO85mN2FlhGAbY0gHEkzA",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/I9_xOTF3DiE/mqdefault_6s.webp?du=3000&sqp=CNr2m4YG&rs=AOn4CLBdlzy786g5kRB4xtOv2L6lwKfHdQ",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=I9_xOTF3DiE"
        },
        {
          "id": "43VtW9BWvxg",
          "title": "Doctor Neiman - Wait For Me (Feat. Micah Martin) [NCS Release]",
          "published_time": "3 weeks ago",
          "duration": 216,
          "view_count": 680252,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/43VtW9BWvxg/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBcUgFMVLRHHsKD2d9oBdTxJ09oSg",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/43VtW9BWvxg/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDYbhGwFEk91kkJRUfjmsGm4huGow",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/43VtW9BWvxg/mqdefault_6s.webp?du=3000&sqp=CKC6nIYG&rs=AOn4CLBtOiCVbl2MsEZLv34y2_tM5ZfMyA",
            "height": 180,
            "width": 320
          },
          "description": "NCS Arcade → our official second channel has just launched, including our brand new 24/7 Livestream! Join in here: ...",
          "channel": {
            "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
            "title": "NoCopyrightSounds",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
          },
          "link": "https://www.youtube.com/watch?v=43VtW9BWvxg"
        }
      ]
    },
    {
      "title": "People also watched",
      "items": [
        {
          "id": "IvVgQHZtpQE",
          "title": "Best of Elektronomia 2018 | Top 20 Songs of Elektronomia",
          "published_time": "2 years ago",
          "duration": 4757,
          "view_count": 2526156,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/IvVgQHZtpQE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDMuVnG1n3vFCydcS25IoGFeATqiA",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/IvVgQHZtpQE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLA3EB6Wp6JT-cDSieyTiTqWcFTCVA",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": null,
          "description": "Elektronomia Top 20 | Best of Elektronomia 2018 Please help Musicbot reach 100k subscriber by join at ...",
          "channel": {
            "id": "UCDsbBjIl0lYZB4IokDLyWIQ",
            "title": "Musicbot",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnjVrECxpse-tll2__nfmWWSM99W2Iyv4o-ksiem6Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCDsbBjIl0lYZB4IokDLyWIQ"
          },
          "link": "https://www.youtube.com/watch?v=IvVgQHZtpQE"
        },
        {
          "id": "_wo5C9qh4xE",
          "title": "How to Use Instagram - Beginner's Guide",
          "published_time": "1 year ago",
          "duration": 1035,
          "view_count": 1229933,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/_wo5C9qh4xE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLB_J9SS_OD7wkKA5COdNPUJiLQZig",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/_wo5C9qh4xE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAlTsnh8uxTNjBSppamUdgHSn89BA",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/_wo5C9qh4xE/mqdefault_6s.webp?du=3000&sqp=CKmenIYG&rs=AOn4CLAHXl4prAXjGDrq44wgCCu_ockh6Q",
            "height": 180,
            "width": 320
          },
          "description": "New to Instagram? Not sure if you are getting the most out of the app? This video will walk you through using Instagram, step by ...",
          "channel": {
            "id": "UCrSvDunJEc1CME4-KvhW_3Q",
            "title": "Howfinity",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnjT1utCm7yjjhEUjPVwBMdwQK8m4G2UYrtP3RoiEw=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCrSvDunJEc1CME4-KvhW_3Q"
          },
          "link": "https://www.youtube.com/watch?v=_wo5C9qh4xE"
        },
        {
          "id": "mt2hJr4ulLU",
          "title": "🔥 Top 50 NoCopyRightSounds | Best of NCS | Most viewed ! Gaming Music | The Best of All Time | 2020",
          "published_time": "11 months ago",
          "duration": 11286,
          "view_count": 2093018,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/mt2hJr4ulLU/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBRHb79Nc1CJFFDdRgecoEyrgrI5A",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/mt2hJr4ulLU/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD_6CNVzwjKoj6XBpCeU7bQr_2G9A",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/mt2hJr4ulLU/mqdefault_6s.webp?du=3000&sqp=CLGanIYG&rs=AOn4CLAoo6XL5iJyVcPg8ib17LBjo9dZ3g",
            "height": 180,
            "width": 320
          },
          "description": "Best of NCS | The Best of All Time Track list: 01. 0:00 Alan Walker - Fade 02. 4:20 Alan Walker - Spectre 03. 8:07 Different Heaven ...",
          "channel": {
            "id": "UCohd6FLJfy3JMmYWJcn4yYA",
            "title": "Phan Trieu",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnjowZb1FENbS5G_vzIIYkktBYUJPhI3g7E_u54Erw=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCohd6FLJfy3JMmYWJcn4yYA"
          },
          "link": "https://www.youtube.com/watch?v=mt2hJr4ulLU"
        },
        {
          "id": "triXo_xCqms",
          "title": "TOP 100 NoCopyrightSounds | Best Of NCS, 6H NoCopyRightSounds | TOP 100 NCS, The Best of all time",
          "published_time": "3 years ago",
          "duration": 22664,
          "view_count": 6625733,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/triXo_xCqms/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCCdwpnnd68xW-Yg-9oz-L3kXWutg",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/triXo_xCqms/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD5VNhq765byHhvrqCMm7vh-cx4dQ",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": null,
          "description": "TOP 100 NoCopyrightSounds | Best Of NCS | 6H NoCopyRightSounds | The Best of all time NCS Song: Best of all time Please ...",
          "channel": {
            "id": "UCDsbBjIl0lYZB4IokDLyWIQ",
            "title": "Musicbot",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnjVrECxpse-tll2__nfmWWSM99W2Iyv4o-ksiem6Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCDsbBjIl0lYZB4IokDLyWIQ"
          },
          "link": "https://www.youtube.com/watch?v=triXo_xCqms"
        },
        {
          "id": "EYKyhCkvsTI",
          "title": "Music Mix 2021 🎧 Remixes of Popular Songs 🎧 EDM Best Music Mix",
          "published_time": "1 week ago",
          "duration": 4530,
          "view_count": 1506623,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/EYKyhCkvsTI/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD5i2NP5JY3bs7Wt0akAsPqoIvpwA",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/EYKyhCkvsTI/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDjxVTgJO8ICWW3Fx-mOjNFA-_TgQ",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/EYKyhCkvsTI/mqdefault_6s.webp?du=3000&sqp=CIqUnIYG&rs=AOn4CLDk8FRzvQMKb9wboFoe5wk_Nk8z0g",
            "height": 180,
            "width": 320
          },
          "description": "Music Mix 2021 Remixes of Popular Songs EDM Best Music Mix Tracklist: 0:00 DJSM, Robbe, New Beat Order - Salt (ft.",
          "channel": {
            "id": "UCp6_KuNhT0kcFk-jXw9Tivg",
            "title": "Magic Music",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnh3uWcwORUsZeEHyV6VEYPQSRGeHSq0haJ8dJyCZg=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCp6_KuNhT0kcFk-jXw9Tivg"
          },
          "link": "https://www.youtube.com/watch?v=EYKyhCkvsTI"
        },
        {
          "id": "xu8rh9Ref4Y",
          "title": "How to Use Facebook - Complete Beginner's Guide",
          "published_time": "9 months ago",
          "duration": 776,
          "view_count": 127527,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/xu8rh9Ref4Y/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLAEdImW27UH7gSbG1m66CaV7qscyg",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/xu8rh9Ref4Y/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBjPbQeX--Za-K7zyxiHZXkaqN9mg",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/xu8rh9Ref4Y/mqdefault_6s.webp?du=3000&sqp=CNSvnIYG&rs=AOn4CLDoAu8k6BfieKPS7T6j48Eknz5hmQ",
            "height": 180,
            "width": 320
          },
          "description": "New To Facebook or haven't been on Facebook for a while? In this video, I'll show you how to use Facebook step by step, from ...",
          "channel": {
            "id": "UCrSvDunJEc1CME4-KvhW_3Q",
            "title": "Howfinity",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnjT1utCm7yjjhEUjPVwBMdwQK8m4G2UYrtP3RoiEw=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCrSvDunJEc1CME4-KvhW_3Q"
          },
          "link": "https://www.youtube.com/watch?v=xu8rh9Ref4Y"
        },
        {
          "id": "V-jwY7uWoP4",
          "title": "Top 20 canciones de NCS",
          "published_time": "2 years ago",
          "duration": 4446,
          "view_count": 5525024,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/V-jwY7uWoP4/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDZYJhEmNsa3ARbn9O--U-bFBortg",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/V-jwY7uWoP4/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDWTfzPftygBhjQcCrSDsLSbIUvxA",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/V-jwY7uWoP4/mqdefault_6s.webp?du=3000&sqp=CNyknIYG&rs=AOn4CLAgJXWPBMhpiYdLVu5XRw3uQMl5dw",
            "height": 180,
            "width": 320
          },
          "description": "Estas son para mi las mejores de el canal... queréis que suba lo misma música pero de otro canal??? Vosotros solo decírmelo!!!",
          "channel": {
            "id": "UCev1VyGGPNwcGBqvPC6yexw",
            "title": "Pelotosino 2945",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnidgpCcZvoHiO1US2qUGIlLekZGd8XFxvDLKPzebg=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCev1VyGGPNwcGBqvPC6yexw"
          },
          "link": "https://www.youtube.com/watch?v=V-jwY7uWoP4"
        },
        {
          "id": "W9iUh23Xrsg",
          "title": "New Year Music Mix 2021 ♫ Best Music 2020 Party Mix ♫ Remixes of Popular Songs",
          "published_time": "5 months ago",
          "duration": 4723,
          "view_count": 37612577,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/W9iUh23Xrsg/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDT1IX4m7jU860Rfcih_N_C0jd_sQ",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/W9iUh23Xrsg/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDxHXa0nrdedqdwjHczYwl3nhSnaQ",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/W9iUh23Xrsg/mqdefault_6s.webp?du=3000&sqp=CPi9nIYG&rs=AOn4CLAN9yfE4FKVhyGpAW3qWvD_b4aBGg",
            "height": 180,
            "width": 320
          },
          "description": "Tracklist: 00:00 Coopex & Yohan Gerber - Radioactive (ft. LUNIS) 02:16 Robbe, New Beat Order & Britt Lari - Kings & Queens ...",
          "channel": {
            "id": "UCp6_KuNhT0kcFk-jXw9Tivg",
            "title": "Magic Music",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnh3uWcwORUsZeEHyV6VEYPQSRGeHSq0haJ8dJyCZg=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCp6_KuNhT0kcFk-jXw9Tivg"
          },
          "link": "https://www.youtube.com/watch?v=W9iUh23Xrsg"
        },
        {
          "id": "YQniVGFuhg8",
          "title": "Storyblocks Review: Is It Worth It In 2021 | IMHO Reviews",
          "published_time": "Streamed 1 year ago",
          "duration": 602,
          "view_count": 33505,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/YQniVGFuhg8/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD1SNRzbjxE-Up9SmkLj0LVbRmcmQ",
              "height": 270,
              "width": 480
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/YQniVGFuhg8/mqdefault_6s.webp?du=3000&sqp=CLianIYG&rs=AOn4CLCMeStm3-40MdBXxAcU9C13bUAYlA",
            "height": 180,
            "width": 320
          },
          "description": "Keywords: Soryblocks review Is Storyblocks worth it? Storyblocks reviews Imho Reviews Affiliate Disclaimer: If you decide to ...",
          "channel": {
            "id": "UCxzWfTJEoVgN_I-_nHTyIlg",
            "title": "Vitaliy Lano IMHO Reviews",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwngG3GrUthVRQO5GFU8ZOtrIb0ZI9aY_jefAsN1q1Q=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UCxzWfTJEoVgN_I-_nHTyIlg"
          },
          "link": "https://www.youtube.com/watch?v=YQniVGFuhg8"
        },
        {
          "id": "ODfy2YIKS1M",
          "title": "ENVATO ELEMENTS | What You Get for $16.50 USD/month",
          "published_time": "1 year ago",
          "duration": 314,
          "view_count": 33889,
          "thumbnails": [
            {
              "url": "https://i.ytimg.com/vi/ODfy2YIKS1M/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCPwGlKYPktj0HGiiyhzoTXUyQ5oA",
              "height": 202,
              "width": 360
            },
            {
              "url": "https://i.ytimg.com/vi/ODfy2YIKS1M/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDWn5ShNngJtaY779U6VCGUxd6q0A",
              "height": 404,
              "width": 720
            }
          ],
          "rich_thumbnail": {
            "url": "https://i.ytimg.com/an_webp/ODfy2YIKS1M/mqdefault_6s.webp?du=3000&sqp=CM22nIYG&rs=AOn4CLBcZXBsix-M3O8Sb1O6s8DXlWMChg",
            "height": 180,
            "width": 320
          },
          "description": "A must have subscription for any type of creator. Envato Elements includes over 740000 stock video clips, over 1000000 still ...",
          "channel": {
            "id": "UC5J-7sW8cVzCGyKN8ym4QsA",
            "title": "Serge M",
            "thumbnails": [
              {
                "url": "https://yt3.ggpht.com/ytc/AAUvwnh3GwC3sFFGyytBNqwGmDxeZZSKZ5CD6qdHrJ9wtg=s68-c-k-c0x00ffffff-no-rj",
                "height": 68,
                "width": 68
              }
            ],
            "link": "https://www.youtube.com/channel/UC5J-7sW8cVzCGyKN8ym4QsA"
          },
          "link": "https://www.youtube.com/watch?v=ODfy2YIKS1M"
        }
      ]
    }
  ],
  "suggestions": [
    "ncs music bass boosted",
    "ncs instrumental",
    "ncs tobu",
    "ncs elektronomia",
    "ncs background music",
    "ncs dj",
    "ncs mortals",
    "ncs fearless",
    "ncs fade",
    "ncs invincible",
    "no copyright song",
    "ncs heart afire",
    "ncs 1 hour",
    "ncs heroes tonight",
    "ncs ringtone",
    "ncs superhero",
    "ncs vlog music",
    "ncs spectre"
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
	search := ytsearch.VideoSearch("ncs")
	search.Limit = 2

	// search.SearchFilter = ytsearch.VideoFilter
	search.SortOrder = ytsearch.ViewCountOrder

	result, err := search.Next()
	if err != nil {
		panic(err)
	}

	jsonstr, _ := json.Marshal(result.Videos)
	fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
[
  {
    "id": "bM7SZ5SBzyY",
    "title": "Alan Walker - Fade [NCS Release]",
    "published_time": "6 years ago",
    "duration": 261,
    "view_count": 439970797,
    "thumbnails": [
      {
        "url": "https://i.ytimg.com/vi/bM7SZ5SBzyY/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCndJeieOK_hJwZkKKLRI4mm7uHtA",
        "height": 270,
        "width": 480
      }
    ],
    "rich_thumbnail": {
      "url": "https://i.ytimg.com/an_webp/bM7SZ5SBzyY/mqdefault_6s.webp?du=3000&sqp=CPK8nIYG&rs=AOn4CLAXmp_ShLdTT517z0exP45P2VaWvA",
      "height": 180,
      "width": 320
    },
    "description": "NCS: Music Without Limitations NCS Spotify: http://spoti.fi/NCS Alan Walker - Fade is included in our debut compilation NCS: ...",
    "channel": {
      "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
      "title": "NoCopyrightSounds",
      "thumbnails": [
        {
          "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
          "height": 68,
          "width": 68
        }
      ],
      "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
    },
    "link": "https://www.youtube.com/watch?v=bM7SZ5SBzyY"
  },
  {
    "id": "K4DyBUG242c",
    "title": "Cartoon - On & On (feat. Daniel Levi) [NCS Release]",
    "published_time": "5 years ago",
    "duration": 208,
    "view_count": 418877980,
    "thumbnails": [
      {
        "url": "https://i.ytimg.com/vi/K4DyBUG242c/hqdefault.jpg?sqp=-oaymwEcCOADEI4CSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCVDGZSQEPvPuF1QyELmT2FV6vdaQ",
        "height": 270,
        "width": 480
      }
    ],
    "rich_thumbnail": {
      "url": "https://i.ytimg.com/an_webp/K4DyBUG242c/mqdefault_6s.webp?du=3000&sqp=CLSxnIYG&rs=AOn4CLD09OYGOdDGp81ZNimc0K3cGVX_hA",
      "height": 180,
      "width": 320
    },
    "description": "Lyrics: Cartoon - On & On feat. Daniel Levi [Verse 1] Hold me close 'til I get up Time is barely on our side I don't wanna waste ...",
    "channel": {
      "id": "UC_aEa8K-EOJ3D6gOs7HcyNg",
      "title": "NoCopyrightSounds",
      "thumbnails": [
        {
          "url": "https://yt3.ggpht.com/ytc/AAUvwnhFuyxDVb5Ls5HDKyBdydj39h9pBkx5I1VhQ7UL5Q=s68-c-k-c0x00ffffff-no-rj",
          "height": 68,
          "width": 68
        }
      ],
      "link": "https://www.youtube.com/channel/UC_aEa8K-EOJ3D6gOs7HcyNg"
    },
    "link": "https://www.youtube.com/watch?v=K4DyBUG242c"
  }
]
```
</details>

## Information

Pull requests are welcome. I'm relatively new to Golang, and I want to see the best practices used in this project.

Only the main functions have been ported. There is a long way to go.

### TODO:
- Better exceptions handling
- Add example file
- Add tests
- Add missing features from youtube-search-python (video & playlist information, search suggestions)

## License

**MIT License**, see [LICENSE](./LICENSE) file for additional information.