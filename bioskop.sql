PGDMP      +        
        }            bioskop    17.4    17.4                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false                       1262    16388    bioskop    DATABASE     i   CREATE DATABASE bioskop WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE bioskop;
                     postgres    false            �            1259    16390    bioskop    TABLE     �   CREATE TABLE public.bioskop (
    id integer NOT NULL,
    nama character varying(100) NOT NULL,
    lokasi character varying(100) NOT NULL,
    rating double precision
);
    DROP TABLE public.bioskop;
       public         heap r       postgres    false            �            1259    16389    bioskop_id_seq    SEQUENCE     �   CREATE SEQUENCE public.bioskop_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.bioskop_id_seq;
       public               postgres    false    218                       0    0    bioskop_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.bioskop_id_seq OWNED BY public.bioskop.id;
          public               postgres    false    217            z           2604    16393 
   bioskop id    DEFAULT     h   ALTER TABLE ONLY public.bioskop ALTER COLUMN id SET DEFAULT nextval('public.bioskop_id_seq'::regclass);
 9   ALTER TABLE public.bioskop ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    217    218    218                      0    16390    bioskop 
   TABLE DATA           ;   COPY public.bioskop (id, nama, lokasi, rating) FROM stdin;
    public               postgres    false    218   �
                  0    0    bioskop_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.bioskop_id_seq', 4, true);
          public               postgres    false    217            |           2606    16395    bioskop bioskop_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.bioskop
    ADD CONSTRAINT bioskop_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.bioskop DROP CONSTRAINT bioskop_pkey;
       public                 postgres    false    218               <   x�3�L�K�T.-H-�M���J�N,*I�440�2���)�*8�����L�b���� Xl%     