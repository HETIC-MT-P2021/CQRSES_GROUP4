# Add command and event

In this example, we want to update title of an article.
Follow this guide step by step to create your own command and event.

**Steps :**
- Command
- Bus
- Event

## Command
Create new route in `ApplyRoutes()` /cmd/v1/article/article.go
```go
r.PUT("/articles/:aggregate_article_id/title", jwt_auth.Operator(jwtAuth), UpdateArticleTitle)
```

Create associated method for this route in /cmd/v1/article/article.ctrl.go
```go
// UpdateArticleTitle will generate a command UpdateArticleTitleCommand
func UpdateArticleTitle(c *gin.Context) {
	aggregateArticleID := c.Param("aggregate_article_id")

	var req database.Article
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	command := cqrs.NewCommandImpl(&commands.UpdateArticleTitleCommand{
		AggregateArticleID: aggregateArticleID,
		Title:              req.Title,
	})

	err := domain.CommandBus.Dispatch(command)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status": "updated",
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": 0,
			"error":  err.Error(),
		})
	}
}
```

Create command in /domain/commands/command.go
```go
// UpdateArticleTitleCommand Command to update title of an article
type UpdateArticleTitleCommand struct {
	Title string
	AggregateArticleID string

}
```

Create handler associated with created command in /domain/commands/handler.go

```go
// UpdateArticleTitleCommandHandler associated to UpdateArticleTitleCommand
type UpdateArticleTitleCommandHandler struct{}

// Handle update title of an article
func (cHandler UpdateArticleTitleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *UpdateArticleTitleCommand:
		message := rabbit.ConsumeMessage{
			EventType: events.ArticleUpdatedTitleEventType,
			Payload: events.ArticleUpdatedTitleEvent{
				AggregateArticleID: cmd.AggregateArticleID,
				Title:       cmd.Title,
			},
		}

		return rabbit.QueueConnector(message)
	default:
		return errors.New("bad command type")
	}
}

// UpdateArticleTitleCommandHandler Creates an instance
func NewUpdateArticleTitleCommandHandler() *UpdateArticleTitleCommandHandler {
	return &UpdateArticleTitleCommandHandler{}
}
```

## Bus
Add mapping between command and handler in `initCommandBus()` /domain/bus.go

```go
_ = CommandBus.AddHandler(commands.NewUpdateArticleTitleCommandHandler(), &commands.UpdateArticleTitleCommand{})
```

Add mapping between event and handler in `initEventBus()` /domain/bus.go

```go
_ = EventBus.AddHandler(events.NewArticleUpdatedTitleEventHandler(), events.ArticleUpdatedTitleEventType)
```

## Event
Add mapping between command and handler in `initCommandBus()` /domain/bus.go

Create event in /domain/events/events.go

```go
//ArticleUpdatedTitleEventType is an event
var ArticleUpdatedTitleEventType = "ArticleUpdatedTitleEvent"

//ArticleUpdatedTitleEvent Event to update title of an article
type ArticleUpdatedTitleEvent struct {
	Title string `json:"title"`
	AggregateArticleID string `json:"aggregate_article_id"`
}
```

Create handler associated with event in /domain/events/handler.go

```go
// ArticleUpdatedTitleEventHandler allows to update an article
type ArticleUpdatedTitleEventHandler struct{}

// Handle Updates a article title
func (eHandler ArticleUpdatedTitleEventHandler) Handle(ev event.Event) error {
	switch evType := ev.Type(); evType {
	case ArticleUpdatedTitleEventType:
		event := ArticleUpdatedTitleEvent{}
		return event.Apply(ev)
	default:
		return errors.New("bad event")
	}
}

// NewArticleUpdatedTitleEventHandler Creates an instance
func NewArticleUpdatedTitleEventHandler() *ArticleUpdatedTitleEventHandler {
	return &ArticleUpdatedTitleEventHandler{}
}
```

Create `Apply()` in /domain/events/apply.go

```go
//Apply To update an aggregate in read-model
//1. Get aggregate from elastic-search
//2. update article state
//3. Update to elastic-search
func (event ArticleUpdatedTitleEvent) Apply(ev event.Event) error {
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	event.AggregateArticleID = payloadMapped["aggregate_article_id"].(string)

	articleFromElastic, err := event.update(payloadMapped)
	if err != nil {
		return err
	}

	if ev.ShouldBeStored() {
		event.storeEventToElastic(articleFromElastic)
	}

	return event.storeReadModel(articleFromElastic)
}
```

Create all methods needed by `Apply()` in /domain/events/processor.go

```go
//------------------------------------------------------------------
// ArticleUpdatedTitleEvent
//------------------------------------------------------------------

//update article state
//@see Action interface
func (event ArticleUpdatedTitleEvent) update(articlePayload map[string]interface{}) (database.Article, error) {
	article, err := elasticsearch.GetReadmodel(event.AggregateArticleID)
	if err != nil {
		return database.Article{}, err
	}

	article.Title = articlePayload["title"].(string)

	return article, nil
}

//storeReadModel An article in db
//@see Action interface
func (event ArticleUpdatedTitleEvent) storeReadModel(article database.Article) error {
	return elasticsearch.StoreReadmodel(article)
}

//storeEventToElastic in db
//@see Action interface
func (event ArticleUpdatedTitleEvent) storeEventToElastic(article database.Article) error {
	createdAt := strconv.FormatInt(time.Now().Unix(), 10)
	newEvent := database.Event{
		ID:        uuid.NewV4().String(),
		EventType: ArticleUpdatedTitleEventType,
		CreatedAt: createdAt,
		Payload:   article,
	}

	return elasticsearch.StoreEvent(newEvent)
}

```