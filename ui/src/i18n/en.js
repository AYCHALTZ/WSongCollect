import deepmerge from 'deepmerge'
import englishMessages from 'ra-language-english'

export default deepmerge(englishMessages, {
  resources: {
    song: {
      name: 'Song |||| Songs',
      fields: {
        albumArtist: 'Album Artist',
        duration: 'Time',
        trackNumber: 'Track #',
        playCount: 'Plays'
      },
      bulk: {
        addToQueue: 'Play Later'
      }
    },
    album: {
      fields: {
        albumArtist: 'Album Artist',
        duration: 'Time',
        songCount: 'Songs',
        playCount: 'Plays'
      },
      actions: {
        playAll: 'Play',
        playNext: 'Play Next',
        addToQueue: 'Play Later',
        shuffle: 'Shuffle'
      }
    }
  },
  ra: {
    auth: {
      welcome1: 'Thanks for installing Navidrome!',
      welcome2: 'To start, create an admin user',
      confirmPassword: 'Confirm Password',
      buttonCreateAdmin: 'Create Admin'
    },
    validation: {
      invalidChars: 'Please only use letter and numbers',
      passwordDoesNotMatch: 'Password does not match'
    }
  },
  menu: {
    library: 'Library',
    settings: 'Settings',
    personal: 'Personal',
    version: 'Version %{version}',
    theme: 'Theme'
  },
  player: {
    panelTitle: 'Play Queue',
    playModeText: {
      order: 'In order',
      orderLoop: 'Repeat',
      singleLoop: 'Repeat One',
      shufflePlay: 'Shuffle'
    }
  }
})
