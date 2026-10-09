┌──────────────────┐
│      TENANT      │
├──────────────────┤
│ PK id            │
│ name             │
│ address          │
│ email            │
│ phone            │
│ status           │
└────────┬─────────┘
         │ 1
         │
         │ N
┌────────▼─────────┐
│       USER       │
├──────────────────┤
│ PK id            │
│ tenant_id FK     │
│ name             │
│ email            │
│ phone            │
│ password_hash    │
│ role             │
│ status           │
└──────┬─────┬─────┘
       │     │
       │     │
       │     └─────────────────────┐
       │                           │
       │ N                         │ N
       │                           │
┌──────▼───────────┐       ┌───────▼──────────┐
│      MEALS       │       │     DEPOSIT      │
├──────────────────┤       ├──────────────────┤
│ PK id            │       │ PK id            │
│ user_id FK       │       │ user_id FK       │
│ tenant_id FK     │       │ tenant_id FK     │
│ date             │       │ amount           │
│ breakfast        │       │ payment_date     │
│ lunch            │       │ method           │
│ dinner           │       │ status           │
└──────────────────┘       └──────────────────┘


┌──────────────────┐
│      BAZAR       │
├──────────────────┤
│ PK id            │
│ tenant_id FK     │
│ user_id FK       │
│ item_name        │
│ quantity         │
│ amount           │
│ purchase_date    │
│ description      │
└────────┬─────────┘
         │
         │ N
         │
         │ 1
┌────────▼─────────┐
│      PAYMENT     │
├──────────────────┤
│ PK id            │
│ user_id FK       │
│ tenant_id FK     │
│ amount           │
│ type             │
│ payment_date     │
│ status           │
└──────────────────┘


┌──────────────────┐
│      NOTICE      │
├──────────────────┤
│ PK id            │
│ tenant_id FK     │
│ created_by FK    │
│ title            │
│ description      │
│ created_at       │
└──────────────────┘


┌──────────────────┐
│      SUBJECT     │
├──────────────────┤
│ PK id            │
│ tenant_id FK     │
│ name             │
│ description      │
│ status           │
└──────────────────┘

## Main relationships
TENANT
  │
  ├──< USERS
  │      │
  │      ├──< MEALS
  │      ├──< DEPOSITS
  │      ├──< PAYMENTS
  │      └──< BAZAR
  │
  ├──< NOTICES
  │
  └──< SUBJECTS

### For this project, I would structure the database roughly like this:

| Table      | Purpose                                    |
| ---------- | ------------------------------------------ |
| `tenants`  | Each separate mess                         |
| `users`    | Mess members/admins                        |
| `meals`    | Daily breakfast/lunch/dinner               |
| `deposits` | Member deposits                            |
| `bazars`   | Grocery/bazar expenses                     |
| `payments` | Other payments/transactions                |
| `notices`  | Mess announcements                         |
| `subjects` | Your application's subject-related feature |
