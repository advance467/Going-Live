CREATE TABLE titles
(
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	description VARCHAR
);

INSERT INTO public.titles ("name",description) VALUES
	 ('I Diari della Speziale','Maomao, una giovane farmacista del quartiere a luci rosse della capitale, viene venduta suo malgrado come sguattera alla corte imperiale.'),
	 ('Kowloon Generic Romance','Nei bassifondi della città murata di Kowloon, un insediamento sovrappopolato e in gran parte non governato della regione di Hong Kong, la trentenne Kujirai e il suo collega Kudo lavorano eroicamente come agenti immobiliari.');
