create table if not exists public.useremails
(
    id int8 not null generated always as identity primary key,
    email varchar(100) not null
);

-- drop table public.useremails

create table if not exists public.friends
(
    id int8 not null generated always as identity primary key,
    first_id int8 not null,
    second_id int8 not null,
    constraint firstemail_fk foreign key (first_id) references public.useremails(id),
    constraint secondemail_fk foreign key (second_id) references public.useremails(id)
);

-- drop table public.friends

create table if not exists public.subscriptions
(
    id int8 not null generated always as identity primary key,
    requestor_id int8 not null,
    target_id int8 not null,
    constraint requestid_fk foreign key (requestor_id) references public.useremails(id),
    constraint targetid_fk foreign key (target_id) references public.useremails(id)
);

--drop table public.subscriptions

create table if not exists public.blocks
(
    id int8 not null generated always as identity primary key,
    requestor_id int8 not null,
    target_id int8 not null,
    constraint requestid_fk foreign key (requestor_id) references public.useremails(id),
    constraint targetid_fk foreign key (target_id) references public.useremails(id)
);





