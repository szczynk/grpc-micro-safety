-- update nulled check-out to now where check-in is not null every day at 3:00am (GMT)
SELECT "cron"."schedule_in_database" (
  "checkout-fillup",
  "0 3 * * *",
  $$UPDATE attendances SET check_out = NOW() WHERE check_in IS NOT NULL AND EXTRACT(DAY FROM check_in) = EXTRACT(DAY FROM NOW())-1 AND check_out IS NULL$$,
  "safety",
  "postgres",
  true
)

-- update pending attendance to rejected every day at 3:00am (GMT)
SELECT "cron"."schedule_in_database" (
  "reject-staled-attendance",
  "0 3 * * *",
  $$UPDATE "attendances" SET "status" = 'rejected', "status_at" = NOW(), status_info: 'Request Time Out' JOIN schedules ON attendances.schedule_id = schedules.id WHERE EXTRACT(DAY FROM schedules.date) = EXTRACT(DAY FROM NOW()) AND attendances.status = 'pending'$$,
  "safety",
  "postgres",
  true
)
