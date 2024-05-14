// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "rating", Type: field.TypeFloat64},
		{Name: "description", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "movie_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_movies_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{MoviesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FoodsColumns holds the columns for the "foods" table.
	FoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "image", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// FoodsTable holds the schema information for the "foods" table.
	FoodsTable = &schema.Table{
		Name:       "foods",
		Columns:    FoodsColumns,
		PrimaryKey: []*schema.Column{FoodsColumns[0]},
	}
	// FoodOrderLinesColumns holds the columns for the "food_order_lines" table.
	FoodOrderLinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "food_id", Type: field.TypeUUID},
		{Name: "transaction_id", Type: field.TypeUUID},
	}
	// FoodOrderLinesTable holds the schema information for the "food_order_lines" table.
	FoodOrderLinesTable = &schema.Table{
		Name:       "food_order_lines",
		Columns:    FoodOrderLinesColumns,
		PrimaryKey: []*schema.Column{FoodOrderLinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "food_order_lines_foods_food_order_lines",
				Columns:    []*schema.Column{FoodOrderLinesColumns[4]},
				RefColumns: []*schema.Column{FoodsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "food_order_lines_transactions_food_order_lines",
				Columns:    []*schema.Column{FoodOrderLinesColumns[5]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MoviesColumns holds the columns for the "movies" table.
	MoviesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "genre", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"UPCOMING", "ONGOING", "OVER"}},
		{Name: "language", Type: field.TypeString},
		{Name: "director", Type: field.TypeString},
		{Name: "cast", Type: field.TypeString},
		{Name: "poster", Type: field.TypeString},
		{Name: "rated", Type: field.TypeString},
		{Name: "duration", Type: field.TypeInt},
		{Name: "trailer", Type: field.TypeString},
		{Name: "opening_day", Type: field.TypeTime},
		{Name: "story", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// MoviesTable holds the schema information for the "movies" table.
	MoviesTable = &schema.Table{
		Name:       "movies",
		Columns:    MoviesColumns,
		PrimaryKey: []*schema.Column{MoviesColumns[0]},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "room_number", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "theater_id", Type: field.TypeUUID},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rooms_theaters_rooms",
				Columns:    []*schema.Column{RoomsColumns[4]},
				RefColumns: []*schema.Column{TheatersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SeatsColumns holds the columns for the "seats" table.
	SeatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "seat_number", Type: field.TypeString},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"STANDARD", "DOUBLE"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "room_id", Type: field.TypeUUID},
	}
	// SeatsTable holds the schema information for the "seats" table.
	SeatsTable = &schema.Table{
		Name:       "seats",
		Columns:    SeatsColumns,
		PrimaryKey: []*schema.Column{SeatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "seats_rooms_seats",
				Columns:    []*schema.Column{SeatsColumns[5]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// ShowTimesColumns holds the columns for the "show_times" table.
	ShowTimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "movie_id", Type: field.TypeUUID},
		{Name: "room_id", Type: field.TypeUUID},
	}
	// ShowTimesTable holds the schema information for the "show_times" table.
	ShowTimesTable = &schema.Table{
		Name:       "show_times",
		Columns:    ShowTimesColumns,
		PrimaryKey: []*schema.Column{ShowTimesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "show_times_movies_showTimes",
				Columns:    []*schema.Column{ShowTimesColumns[5]},
				RefColumns: []*schema.Column{MoviesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "show_times_rooms_showTimes",
				Columns:    []*schema.Column{ShowTimesColumns[6]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TheatersColumns holds the columns for the "theaters" table.
	TheatersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "address", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "phone_number", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TheatersTable holds the schema information for the "theaters" table.
	TheatersTable = &schema.Table{
		Name:       "theaters",
		Columns:    TheatersColumns,
		PrimaryKey: []*schema.Column{TheatersColumns[0]},
	}
	// TicketsColumns holds the columns for the "tickets" table.
	TicketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "is_booked", Type: field.TypeBool, Default: false},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "seat_id", Type: field.TypeUUID},
		{Name: "show_time_id", Type: field.TypeUUID},
		{Name: "transaction_id", Type: field.TypeUUID, Nullable: true},
	}
	// TicketsTable holds the schema information for the "tickets" table.
	TicketsTable = &schema.Table{
		Name:       "tickets",
		Columns:    TicketsColumns,
		PrimaryKey: []*schema.Column{TicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickets_seats_tickets",
				Columns:    []*schema.Column{TicketsColumns[5]},
				RefColumns: []*schema.Column{SeatsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tickets_show_times_tickets",
				Columns:    []*schema.Column{TicketsColumns[6]},
				RefColumns: []*schema.Column{ShowTimesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tickets_transactions_tickets",
				Columns:    []*schema.Column{TicketsColumns[7]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "total", Type: field.TypeFloat64},
		{Name: "code", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "PAID", "CANCEL"}, Default: "PENDING"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_users_transactions",
				Columns:    []*schema.Column{TransactionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "displayname", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "password", Type: field.TypeString, Size: 255},
		{Name: "is_locked", Type: field.TypeBool, Default: false},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"CUSTOMER", "STAFF", "TICKET_MANAGER", "ADMIN"}, Default: "CUSTOMER"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		FoodsTable,
		FoodOrderLinesTable,
		MoviesTable,
		RoomsTable,
		SeatsTable,
		SessionsTable,
		ShowTimesTable,
		TheatersTable,
		TicketsTable,
		TransactionsTable,
		UsersTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = MoviesTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	FoodOrderLinesTable.ForeignKeys[0].RefTable = FoodsTable
	FoodOrderLinesTable.ForeignKeys[1].RefTable = TransactionsTable
	RoomsTable.ForeignKeys[0].RefTable = TheatersTable
	SeatsTable.ForeignKeys[0].RefTable = RoomsTable
	ShowTimesTable.ForeignKeys[0].RefTable = MoviesTable
	ShowTimesTable.ForeignKeys[1].RefTable = RoomsTable
	TicketsTable.ForeignKeys[0].RefTable = SeatsTable
	TicketsTable.ForeignKeys[1].RefTable = ShowTimesTable
	TicketsTable.ForeignKeys[2].RefTable = TransactionsTable
	TransactionsTable.ForeignKeys[0].RefTable = UsersTable
}
