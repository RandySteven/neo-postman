package queries

const (
	GetExpectedAndUnexpectedDataQuery GoQuery = `
		SELECT
		SUM(CASE WHEN result_status = 1 THEN 1 ELSE 0 END) AS expected,
		SUM(CASE WHEN result_status = 2 THEN 1 ELSE 0 END) AS unexpected
	FROM
		test_datas
	WHERE
		result_status IN (1, 2);
	`

	GetAvgResponseTimePerAPIQuery GoQuery = `
		SELECT uri, AVG(response_time) as avg_time
			FROM test_datas
		GROUP BY uri 
	`
)
