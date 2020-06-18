using System.Collections.Generic;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace.Configuration;
using OpenTelemetry.Trace.Samplers;

namespace temperature
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers();
            services.AddOpenTelemetry((builder) => builder
                .UseJaeger(o =>
                {
                    o.ServiceName = this.Configuration.GetValue<string>("Jaeger:ServiceName");
                    o.AgentHost = this.Configuration.GetValue<string>("Jaeger:Host");
                    o.AgentPort = this.Configuration.GetValue<int>("Jaeger:Port");
                })
                .SetSampler(new AlwaysSampleSampler())
                .AddDependencyCollector(config =>
                {
                    config.SetHttpFlavor = true;
                    
                })
                .AddRequestCollector()
                .SetResource(new Resource(new Dictionary<string, object>
                {
                    { "service.name", "temperature" }
                })));
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseRouting();

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
