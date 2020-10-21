/* this is an example of a linked list */

package main

import "fmt"

type song struct {
	name   string
	artist string
	next   *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) error {

	s := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = s
	}
	return nil
}

func (p *playlist) showAllSongs() error {

	currentNode := p.head

	if currentNode == nil {
		fmt.Println("Playlist is empty")
		return nil
	}

	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

func (p *playlist) startBeginningofSong() *song {
	p.nowPlaying = p.head

	return p.nowPlaying
}

func (p *playlist) nextSong() *song {
	p.nowPlaying = p.nowPlaying.next

	return p.nowPlaying
}

func main() {
	newPlaylist := createPlaylist("warren's jam")

	newPlaylist.addSong("Mama Mia", "Idontknow")
	newPlaylist.addSong("Who let the dogs out", "Bahaman")

	newPlaylist.showAllSongs()
}
