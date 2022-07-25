import mock from '../MockConfig';
import analytics from '../../db/dashboard/analytics';
import ecommerce from '../../db/dashboard/ecommerce';
import crm from '../../db/dashboard/crm';
import crypto from '../../db/dashboard/crypto';
import metrics from '../../db/dashboard/metrics';
import widgets from '../../db/dashboard/widgets';
import healthCare from '../../db/dashboard/healthCare';
import academy from '../../db/dashboard/academy';

// Define all mocks of dashboard
mock.onGet("/dashboard/analytics").reply(200, analytics);

mock.onGet("/dashboard/ecommerce").reply(200, ecommerce);

mock.onGet("/dashboard/crm").reply(200, crm);

mock.onGet("/dashboard/crypto").reply(200, crypto);

mock.onGet("/dashboard/health_care").reply(200, healthCare);

mock.onGet("/dashboard/academy").reply(200, academy);

mock.onGet("/dashboard/metrics").reply(200, metrics);

mock.onGet("/dashboard/widgets").reply(200, widgets);
