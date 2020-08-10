create table urls
(
   	id uuid default uuid_generate_v4(),
    val varchar not null,
	primary key(id)
);
create table trackers
(
	id uuid default uuid_generate_v4(),
	name varchar not null,
	selector varchar not null,
	primary key(id)
);
create table url_tracker
(
    url_id uuid,
	tracker_id uuid,
	constraint fk_url
		foreign key (url_id)
			references urls(id),
	constraint fk_tracker
		foreign key (tracker_id)
			references trackers(id)
);