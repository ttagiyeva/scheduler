package repository

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	"github.com/ttagiyeva/scheduler/internal/config"
	"github.com/ttagiyeva/scheduler/internal/scheduler/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

type FirestoreRepo struct {
	log    *zap.SugaredLogger
	config *config.Config
	client *firestore.Client
}

//NewFirestoreRepo creates Firestore Client and FirestoreRepo instance
func NewFirestoreRepo(lc fx.Lifecycle, log *zap.SugaredLogger, config *config.Config) (*FirestoreRepo, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.FirestoreConfig.ProjectName)

	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return &FirestoreRepo{
		log:    log,
		config: config,
		client: client}, nil
}

//Save creates a scheduler document in forestore collection
func (f *FirestoreRepo) Save(ctx context.Context, s *domain.Scheduler) error {

	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(s.OrderName).Set(ctx, s)
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}

//Get reterieves a scheduler document of given id
func (f *FirestoreRepo) Get(ctx context.Context, orderName string) (*domain.Scheduler, error) {

	doc, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(orderName).Get(ctx)
	if err != nil {
		f.log.Error(err)
		return nil, err
	}

	body, err := json.Marshal(doc.Data())
	if err != nil {
		f.log.Error(err)
		return nil, err
	}

	scheduler := &domain.Scheduler{}
	if err := json.Unmarshal(body, scheduler); err != nil {
		f.log.Error(err)
		return nil, err
	}

	return scheduler, nil
}

//GetAll retrieves queried documents
func (f *FirestoreRepo) GetAll(ctx context.Context, path string, op string, value interface{}) ([]*domain.Scheduler, error) {

	var datas = make([]*domain.Scheduler, 0)
	iter := f.client.Collection(f.config.FirestoreConfig.CollectionName).Where(path, op, value).Documents(ctx)

	defer iter.Stop()

	for {

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			f.log.Error(err)
			return nil, err
		}

		var data *domain.Scheduler
		if err := doc.DataTo(data); err != nil {
			f.log.Error(err)
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

//Update updates drone_name field of a scheduler document
func (f *FirestoreRepo) Update(ctx context.Context, s *domain.Scheduler) error {

	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(s.OrderName).Update(ctx, []firestore.Update{
		{
			Path:  "drone_name",
			Value: s.DroneName,
		},
	})
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}

//Delete deletes document of given id
func (f *FirestoreRepo) Delete(ctx context.Context, orderName string) error {

	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(orderName).Delete(ctx)
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}

/*
func (f *FirestoreRepo) getFields(scheduler domain.Scheduler) []firestore.Update {

	updates := make([]firestore.Update, 0)
	val := reflect.ValueOf(scheduler)

	for i := 0; i < val.Type().NumField(); i++ {

		field := val.Type().Field(i)
		value := reflect.Indirect(val).FieldByName(field.Name)

		if !value.IsZero() {
			update := firestore.Update{
				Path:  field.Tag.Get("firestore"),
				Value: value,
			}
			updates = append(updates, update)
		}

	}

	return updates
}
*/
