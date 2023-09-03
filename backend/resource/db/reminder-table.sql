DROP TABLE IF EXISTS public.reminder;

CREATE TABLE public.reminder
(
    id   smallserial not null,
    task varchar     not null,
    done boolean     not null,
    constraint id_pk primary key ("id")
);

-- default data
-- INSERT INTO public.reminder (task, done) VALUES ('Buy groceries', false);
-- INSERT INTO public.reminder (task, done) VALUES ('Finish work report', false);
-- INSERT INTO public.reminder (task, done) VALUES ('Pay bills', true);
