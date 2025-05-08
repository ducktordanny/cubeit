# CubeIt Planning Notes

## What we want?

### Users

For a first step, I want users to be able to log in through their WCA account.

- [x] [Check WCA API docs for oauth](https://www.worldcubeassociation.org/help/api)

I've created a WCA application for the authentication, which can be checked here: [WCA App](https://www.worldcubeassociation.org/oauth/applications/1317)

`https://www.worldcubeassociation.org/api/v0/me` sends back the following JSON response:

```json
{
  "me": {
    "id": 11004,
    "created_at": "2016-02-06T11:42:15.000Z",
    "updated_at": "2025-04-25T05:30:49.000Z",
    "name": "Dániel Lázár",
    "wca_id": "2014LAZA04",
    "gender": "m",
    "country_iso2": "HU",
    "url": "https://www.worldcubeassociation.org/persons/2014LAZA04",
    "country": {
      "id": "Hungary",
      "name": "Hungary",
      "continent_id": "_Europe",
      "iso2": "HU"
    },
    "delegate_status": null,
    "class": "user",
    "teams": [],
    "avatar": {
      "id": 66689,
      "status": "approved",
      "thumbnail_crop_x": 155,
      "thumbnail_crop_y": 0,
      "thumbnail_crop_w": 1863,
      "thumbnail_crop_h": 1863,
      "url": "https://avatars.worldcubeassociation.org/qs67jk4b477nafspfk5rgjsulhm3",
      "thumb_url": "https://avatars.worldcubeassociation.org/kapxualu4z4x7ncjtobcv6ecdm5x",
      "is_default": false,
      "can_edit_thumbnail": true
    },
    "email": "..."
  }
}
```

(scope is needed to be set to "public email")
What if I want to fetch other existing users?

`https://www.worldcubeassociation.org/api/v0/persons/{wca_id}` can also fetch user data, and returns some additional metadata, too, eg.:

```json
{
  "person": {
    // same as .me
  },
  "competition_count": 10,
  "personal_records": {
    // all event records, it's a lot
  },
  "medals": {
    "gold": 0,
    "silver": 0,
    "bronze": 0,
    "total": 0
  },
  "records": {
    "national": 2,
    "continental": 0,
    "world": 0,
    "total": 1
  }
}
```

This endpoint returns back a lot of useful and nice info, though probably not needed. I don't want to show WCA records and stats on CubeIt, but just have a link to the WCA profile instead.

I also want to store users in my own db, so I can have some additional features, like follow and maybe some additional info later.

### Solves

> The main topic of the site should be to be able to store personal records made at home and share solves with others.

For this, we need:

- Category (333, 222, 444, 555, 666, etc.)
- Time(s)
- Scramble(s)
  - On the frontend, draw the preview of scramble, use cstimer module
- Penalty [Optional]
- Solution [Optional]
- Image [Optional]
  - Probably won't include in first phase
- Likes, comments
  - If any of this happens we could have in app (or maybe browser based) notifications about it
- Maybe some sort of repost? Do a solve on someone's scramble and reference that
  - In the first phase, this should probably be an optional feature, could be added later on, too
- Mark a solve as PR, or have this automatically (always marking the best time as PR)?
  - In case of reposting the solve should not count as PR
- External share? Preview? Kinda like Strava
- Share multiple solves in one post, e.g. ao5, ao12, etc. (allowing above gives me concerns in storage space, PoC needed)
  - In this case, should also think of optimization. In case if we could share bigger averages, like ao100 and above, we should handle data responses in chunks.

Should have a feed:

- Showing followed people's posts
- Optimization: Responses in chunks

Should be able to look at own solves:

- Also needs chunks

On the frontend we could also add some stats

### Roles?

An admin role may be needed?

## Technologies

Frontend:

- Angular
- Angular Material

Backend:

- Go
- Gin

Db:

- PostgreSQL
- Running inside Docker

Other:

- I don't want to use Nx, as it has a lot of overhead of configuring everything, may switch to it later though. However, right now I don't want to spend too much time on configuring apps, libraries, and dependencies, etc.
