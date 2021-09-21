CREATE
	[DEFINER = { user | CURRENT_USER }]
	EVENT
	[IF NOT EXISTS]
	event_name
	ON SCHEDULE
	[ AT timestamp [+ INTERVAL interval] | EVERY interval ... ]
	[ STARTS timestamp [+ INTERVAL interval] ... ]
	[ ENDS timestamp [+ INTERVAL interval] ... ]
	[ ON COMPLETION [NOT] PRESERVE ]
	[ ENABLE | DISABLE | DISABLE ON SLAVE ]
	[ COMMENT 'comment' ]
	DO
	BEGIN

	END

