--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Homebrew)
-- Dumped by pg_dump version 14.9 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: macbookpro
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';


ALTER DATABASE postgres OWNER TO macbookpro;

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: macbookpro
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: disbursement; Type: SCHEMA; Schema: -; Owner: macbookpro
--

CREATE SCHEMA disbursement;


ALTER SCHEMA disbursement OWNER TO macbookpro;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: banks; Type: TABLE; Schema: disbursement; Owner: macbookpro
--

CREATE TABLE disbursement.banks (
    id integer NOT NULL,
    bank_code character varying(25),
    bank_name character varying(100),
    create_ts timestamp without time zone,
    update_ts timestamp without time zone
);


ALTER TABLE disbursement.banks OWNER TO macbookpro;

--
-- Name: banks_id_seq; Type: SEQUENCE; Schema: disbursement; Owner: macbookpro
--

CREATE SEQUENCE disbursement.banks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE disbursement.banks_id_seq OWNER TO macbookpro;

--
-- Name: banks_id_seq; Type: SEQUENCE OWNED BY; Schema: disbursement; Owner: macbookpro
--

ALTER SEQUENCE disbursement.banks_id_seq OWNED BY disbursement.banks.id;


--
-- Name: disbursements; Type: TABLE; Schema: disbursement; Owner: macbookpro
--

CREATE TABLE disbursement.disbursements (
    id integer NOT NULL,
    account_number character varying(25),
    account_name character varying(100),
    amount double precision,
    receipt_email character varying(100),
    remark text,
    ref_number character varying(100),
    external_id character varying(100),
    status character varying(15),
    create_ts timestamp without time zone,
    update_ts timestamp without time zone
);


ALTER TABLE disbursement.disbursements OWNER TO macbookpro;

--
-- Name: disbursements_id_seq; Type: SEQUENCE; Schema: disbursement; Owner: macbookpro
--

CREATE SEQUENCE disbursement.disbursements_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE disbursement.disbursements_id_seq OWNER TO macbookpro;

--
-- Name: disbursements_id_seq; Type: SEQUENCE OWNED BY; Schema: disbursement; Owner: macbookpro
--

ALTER SEQUENCE disbursement.disbursements_id_seq OWNED BY disbursement.disbursements.id;


--
-- Name: banks id; Type: DEFAULT; Schema: disbursement; Owner: macbookpro
--

ALTER TABLE ONLY disbursement.banks ALTER COLUMN id SET DEFAULT nextval('disbursement.banks_id_seq'::regclass);


--
-- Name: disbursements id; Type: DEFAULT; Schema: disbursement; Owner: macbookpro
--

ALTER TABLE ONLY disbursement.disbursements ALTER COLUMN id SET DEFAULT nextval('disbursement.disbursements_id_seq'::regclass);


--
-- Data for Name: banks; Type: TABLE DATA; Schema: disbursement; Owner: macbookpro
--

INSERT INTO disbursement.banks (id, bank_code, bank_name, create_ts, update_ts) VALUES (1, 'BCA', 'BANK BCA', '2024-03-14 10:11:09', '2024-03-14 10:11:11');


--
-- Data for Name: disbursements; Type: TABLE DATA; Schema: disbursement; Owner: macbookpro
--

INSERT INTO disbursement.disbursements (id, account_number, account_name, amount, receipt_email, remark, ref_number, external_id, status, create_ts, update_ts) VALUES (1, '358129497', 'Timothy Corkery', 1000000, 'joni.iskndr92@gmail.com', 'buat jajan', '17BC92F138647860', '', 'FAILED', NULL, NULL);
INSERT INTO disbursement.disbursements (id, account_number, account_name, amount, receipt_email, remark, ref_number, external_id, status, create_ts, update_ts) VALUES (2, '358129497', 'Timothy Corkery', 1000000, 'joni.iskndr92@gmail.com', 'buat jajan', '17BC937D4A0DB040', '', 'FAILED', '2024-03-14 15:15:24.816553', NULL);
INSERT INTO disbursement.disbursements (id, account_number, account_name, amount, receipt_email, remark, ref_number, external_id, status, create_ts, update_ts) VALUES (3, '358129497', 'Timothy Corkery', 1000000, 'joni.iskndr92@gmail.com', 'buat jajan', '17BC939C62825238', '2', 'SUCCESS', '2024-03-14 15:17:38.370847', '2024-03-14 15:40:44.782648');


--
-- Name: banks_id_seq; Type: SEQUENCE SET; Schema: disbursement; Owner: macbookpro
--

SELECT pg_catalog.setval('disbursement.banks_id_seq', 1, true);


--
-- Name: disbursements_id_seq; Type: SEQUENCE SET; Schema: disbursement; Owner: macbookpro
--

SELECT pg_catalog.setval('disbursement.disbursements_id_seq', 3, true);


--
-- Name: banks banks_pkey; Type: CONSTRAINT; Schema: disbursement; Owner: macbookpro
--

ALTER TABLE ONLY disbursement.banks
    ADD CONSTRAINT banks_pkey PRIMARY KEY (id);


--
-- Name: disbursements disbursements_pk; Type: CONSTRAINT; Schema: disbursement; Owner: macbookpro
--

ALTER TABLE ONLY disbursement.disbursements
    ADD CONSTRAINT disbursements_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

