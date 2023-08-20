package repository

import (
	"context"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

func (p *PeopleRepository) CacheCheckNicknameExists(nickname string) (bool, error) {
	ctx := context.Background()
	return p.CacheCon.Do(ctx, p.CacheCon.B().Get().Key("people_nickname_"+nickname).Build()).AsBool()
}
func (p *PeopleRepository) CacheCheckIdExists(id uuid.UUID) (*people.People, error) {
	ctx := context.Background()
	result, err := p.CacheCon.Do(ctx, p.CacheCon.B().Get().Key("people_"+id.String()).Build()).ToString()
	if err != nil {
		return nil, err
	}
	var peopleData *people.People
	err = sonic.UnmarshalString(result, peopleData)
	if err != nil {
		return nil, err
	}
	return peopleData, nil
}
func (p *PeopleRepository) CacheSavePeople(people *people.People) error {
	ctx := context.Background()
	peopleStringData, err := sonic.MarshalString(people)
	if err != nil {
		return err
	}

	savePeople := p.CacheCon.B().Set().Key("people_" + people.Id).Value(peopleStringData).Build()
	saveNickname := p.CacheCon.B().Setbit().Key("people_nickname_" + people.Nickname).Offset(0).Value(0).Build()
	for _, resp := range p.CacheCon.DoMulti(ctx, saveNickname, savePeople) {
		if err := resp.Error(); err != nil {
			return err
		}
	}
	return nil
}
