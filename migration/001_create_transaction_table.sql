-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "transaction" (
    "id" INT  NOT NULL,
    "request_id" INT,
    "terminal_id" INT,
    "partner_object_id" INT,
    "amount_total" FLOAT,
    "amount_original" FLOAT,
    "commission_ps" FLOAT,
    "commission_client" FLOAT,
    "commission_provider" FLOAT,
    "date_input" Timestamp Without Time Zone,
    "date_post" Timestamp Without Time Zone,
    "status" VARCHAR(8),
    "payment_type" VARCHAR(4),
    "payment_number" VARCHAR(10),
    "service_id" INT,
    "service" VARCHAR(17),
    "payee_id" INT,
    "payee_name" VARCHAR(10),
    "payee_bank_mfo" INT,
    "payee_bank_account" VARCHAR(17),
    "payment_narrative" VARCHAR(255),
    PRIMARY KEY ("id"),
    CONSTRAINT "unique_transaction_id" UNIQUE("id")
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "transaction";