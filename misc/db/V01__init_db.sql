-- tbl rekening
CREATE TABLE public.tbl_rekening (
    id serial NOT NULL,
    norek varchar NOT NULL,
    create_date timestamp with time zone NULL,
    update_date timestamp with time zone NULL,
    saldo numeric NOT NULL DEFAULT 0
);

ALTER TABLE public.tbl_rekening ADD CONSTRAINT tbl_rekening_id_pk PRIMARY KEY (id);

-- tbl harga
CREATE TABLE public.tbl_harga (
    id serial4 NOT NULL,
    harga_topup numeric NOT NULL,
    harga_buyback numeric NOT NULL,
    reff_id varchar NOT NULL,
    is_active boolean NOT NULL,
    admin_id varchar NOT NULL,
    create_date timestamptz NULL,
    update_date timestamptz NULL
);

ALTER TABLE public.tbl_harga ADD CONSTRAINT tbl_harga_id_pk PRIMARY KEY (id);

--tbl transaksi
CREATE TABLE public.tbl_transaksi (
    id bigserial NOT NULL,
    rekening_id bigint NOT NULL,
    harga_buyback numeric NOT NULL,
    gram numeric NOT NULL,
    "type" varchar NOT NULL,
    reff_id varchar NOT NULL,
    saldo_akhir numeric NOT NULL,
    transaction_date timestamp with time zone NOT NULL,
    create_date timestamp with time zone NULL,
    update_date timestamp with time zone NULL,
    harga_topup numeric NOT NULL,
);

ALTER TABLE public.tbl_transaksi ADD CONSTRAINT tbl_transaksi_id_pk PRIMARY KEY (id);
ALTER TABLE public.tbl_transaksi ADD CONSTRAINT tbl_transaksi_rek_id_fk FOREIGN KEY (id) REFERENCES public.tbl_rekening(id);

