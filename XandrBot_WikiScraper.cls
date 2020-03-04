public with sharing class XandrBot_WikiScraper {
	public class JobOutput {
		@Invocablevariable(required=true)
		public String sJobPosition;

	}

	@InvocableMethod(label='Get Available Jobs')
	public static List<JobOutput> getAvaialbeJobs() {
		List<JobOutput> jobs = new List<JobOutput>();

		Http http = new Http();
		HttpRequest req = new HttpRequest();
		req.setEndpoint('https://xandrwikiwebscraper.herokuapp.com/');
		req.setMethod('GET');
		HttpResponse res = http.send(req);

		if (res.getStatusCode() == 200) {
			Map<String, Object> data = (Map<String, Object>) JSON.deserializeUntyped(res.getBody());
			for (Object position : (List<Object>) data.get('positions')) {
				JobOutput job = new JobOutput();
				job.sJobPosition = (String) position;
				jobs.add(job);
			}

		}

		return jobs;

	}

}
