type Repo interface {
	SaveUser(user models.User) error
	UserByEmail(email string) (models.User, error)

	GetAllUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)

	InsertTask(task models.Task) error
	UpdateTask(task models.Task) error
}

type Service struct {
	Store Repo
}

func NewService(db Repo) *Service {
	return &Service{
		Store: db,
	}
}
