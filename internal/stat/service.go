package stat

import (
	"go/adv-demo/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (service *StatService) AddClick() {
	for msg := range service.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			linkId, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited data ", msg.Data)
				continue
			} else {
				service.StatRepository.AddClick(linkId)
			}
		}
	}
}
