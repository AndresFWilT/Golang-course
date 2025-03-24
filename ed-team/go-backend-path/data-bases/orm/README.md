# ORM

it let us map the structure of the db in go structs. It has encapsulated the complex logic of the db and will let us use only the exposed methods to make operations in the db. The ORM has the drivers implemented.

For that use `GORM`, look at the [GORM Models Standard](https://gorm.io/docs/models.html)

## Conventions

### Primary Key
GORM uses a field named ID as the default primary key for each model.

### Table Names

By default, GORM converts struct names to snake_case and pluralizes them for table names. For instance, a User struct becomes users in the database, and a GormUserName becomes gorm_user_names.

### Column Names

GORM automatically converts struct field names to snake_case for column names in the database.

### Timestamp Fields

GORM uses fields named CreatedAt and UpdatedAt to automatically track the creation and update times of records.

### Example

It's important to specify if a field is null or not.

```Go
type User struct {
  ID           uint           // Standard field for the primary key
  Name         string         // A regular string field
  Email        *string        // A pointer to a string, allowing for null values
  Age          uint8          // An unsigned 8-bit integer
  Birthday     *time.Time     // A pointer to time.Time, can be null
  MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
  ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
  CreatedAt    time.Time      // Automatically managed by GORM for creation time
  UpdatedAt    time.Time      // Automatically managed by GORM for update time
  ignored      string         // fields that aren't exported are ignored
}
```

### Field level permission

```Go
type User struct {
  Name string `gorm:"<-:create"` // allow read and create
  Name string `gorm:"<-:update"` // allow read and update
  Name string `gorm:"<-"`        // allow read and write (create and update)
  Name string `gorm:"<-:false"`  // allow read, disable write permission
  Name string `gorm:"->"`        // readonly (disable write permission unless it configured)
  Name string `gorm:"->;<-:create"` // allow read and create
  Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
  Name string `gorm:"-"`            // ignore this field when write and read with struct
  Name string `gorm:"-:all"`        // ignore this field when write, read and migrate with struct
  Name string `gorm:"-:migration"`  // ignore this field when migrate with struct
}
```

It's recommended to do the following

```Go
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OurModel struct {
	gorm.Model
	// fields
}
```

If we want to define a different ID:

```Go
type Animals struct {
	AnimalID    int64  `gorm:primary_key`
}
```

By default, GORM creates the table names with plural: `users` and in lower. But we can update the names with this:

```Go
type Animals struct {
    AnimalID    int64
}

func (Animals) TableName() string {
	return "animals"
}
```

The fields are created with snake_case:

```text
createdAt = created_at
```

To override the behavior do:

```Go
type Animals struct {
    AnimalID    int64   `gorm:column:beast_super_id`	
}
``` 

## One to Many

Look at the following sample:

```Go
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignkey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number  string
	UserRefer  uint
}
```

## CRUD in GORM

When we have our models and the storage connection to the db with a singleton pattern, we can:

### Create

To insert into the db via GORM, we need to create a pointer of the model and then pass the dereference var to the Create method from the unique instance of our DB (singleton) connection to the db:

```Go
product1 := model.Product{
		Name:  "Curso de Python",
		Price: 200,
	}
	
	storage.DB().Create(&product1)
```

### GetAll

To get all the elements of a table, we need to create a slice of the model, then use pointers to have access to the data, also with the `Find` method to get all. (use for to look at all the items):

```Go
products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}
```

### Get

To get an specific model by id, we can create an instance of the product, and use `First` Method:

```Go
myProduct := model.Product{}

	storage.DB().First(&myProduct, 2)
	fmt.Println(myProduct)
```

### Update

To update we can use the method `Updates` we need to create an instance:

```Go
myProduct := model.Product{}
	myProduct.ID = 2

	storage.DB().Model(&myProduct).Updates(
		model.Product{Name: "Curso de Java", Price: 120},
	)
```

### Delete

We can delete via GORM in two ways: permanently or make a soft delete (better)

#### Soft Delete

To soft delete we can use the `Delete` method, this will update the `deleted_at` column from the Model of GORM

```Go
myProduct := model.Product{}
	myProduct.ID = 4

	storage.DB().Delete(&myProduct)
```

#### Delete Permanently

To delete permanently we just need to use the `Unscoped().Delete()` method (not recommended)

```Go
myProduct1 := model.Product{}
	myProduct1.ID = 6

	storage.DB().Unscoped().Delete(&myProduct1)
```
