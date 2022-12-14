import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:footer/footer.dart';
import 'package:footer/footer_view.dart';
import 'package:url_launcher/url_launcher.dart';

class AboutUs extends StatelessWidget {
  const AboutUs({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("About Us"),
        centerTitle: true,
        backgroundColor: Colors.blue,
      ),
      // drawer: new Drawer(),
      body: FooterView(
        footer: Footer(
            backgroundColor: Colors.blue,
            padding: const EdgeInsets.all(5.0),
            alignment: Alignment.bottomCenter,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              // mainAxisAlignment: MainAxisAlignment.spaceAround,
              children: <Widget>[
                Center(
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: <Widget>[
                      Container(
                          height: 45.0,
                          width: 45.0,
                          child: Center(
                            child: Card(
                              elevation: 5.0,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(
                                    25.0), // half of height and width of Image
                              ),
                              child: IconButton(
                                icon: const Icon(Icons.email, size: 20.0),
                                color: const Color.fromARGB(255, 23, 45, 84),
                                onPressed: () {},
                              ),
                            ),
                          )),
                      Container(
                          height: 45.0,
                          width: 45.0,
                          child: Center(
                            child: Card(
                              elevation: 5.0,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(
                                    25.0), // half of height and width of Image
                              ),
                              child: IconButton(
                                icon: const Icon(
                                  Icons.contact_support_rounded,
                                  size: 20.0,
                                ),
                                color: const Color.fromARGB(255, 23, 45, 84),
                                onPressed: () {},
                              ),
                            ),
                          )),
                      Container(
                          height: 45.0,
                          width: 45.0,
                          child: Center(
                            child: Card(
                              elevation: 5.0,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(
                                    25.0), // half of height and width of Image
                              ),
                              child: IconButton(
                                icon: const Icon(
                                  Icons.call,
                                  size: 20.0,
                                ),
                                color: const Color.fromARGB(255, 23, 45, 84),
                                onPressed: () {},
                              ),
                            ),
                          )),
                    ],
                  ),
                ),
                const Text(
                  '\nCopyright ??2022, All Rights Reserved.',
                  style: const TextStyle(
                      fontWeight: FontWeight.w300,
                      fontSize: 12.0,
                      color: Color.fromARGB(255, 237, 240, 243)),
                ),
                const Text(
                  'Powered by Yu Long, Diana Lenghel, Ian Foster',
                  style: TextStyle(
                      fontWeight: FontWeight.w300,
                      fontSize: 12.0,
                      color: Color.fromARGB(255, 237, 240, 243)),
                ),
              ],
            )),
        flex: 10,
        children: <Widget>[
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                color: Colors.indigo[50],
                alignment: Alignment.center,
                child: const Text(
                  'The Dublin Bus DIY Journey Planner app collates bus timing and location '
                  'information and presents them as tailor-made, easy-to-read journey plans. '
                  'The app acts as your own personal door-to-door travel planner for trips in Ireland. '
                  'Whether you plan on travelling now or in the near future, '
                  'the DublinBus DIY Journey Planner app allows you to seamlessly plan your trip!'
                  'Our easy to use route planner for Dublin helps guide you to your destination. '
                  'Find the fastest route directions with multiple stops and most convenient travel itinerary using our maps.'
                  '\n'
                  '\n  Step 1: Enter your departure location in the ORIGIN field of the route '
                  'planner.'
                  '\n  Step 2: Insert your destination in the DESTINATION field on the journey '
                  'planner.'
                  '\n  Step 3: Choose the Departure or Arrival time and then click "Plan". ',
                  textAlign: TextAlign.left,
                  style: TextStyle(fontSize: 15),
                ),
              ),
            ],
          ),

          //containers for the links to mobile app and leap card
          Container(
            color: Colors.indigo[50],
            alignment: Alignment.center,
            child: RichText(
              text: TextSpan(
                children: [
                  const TextSpan(
                    text: '''                                   \u{1F4F1}Download the mobile app here\n\n''',
                    style: TextStyle(
                        color: Colors.black, fontWeight: FontWeight.bold, height: 2, fontSize: 30),
                  ),
                  TextSpan(
                      text: '                                   Download Mobile App - iOS',
                      style: const TextStyle(color: Colors.blue),
                      recognizer: TapGestureRecognizer()
                        ..onTap = () {
                          launch('https://dublinbus-diy.site/download/DublinBus.ipa');
                        }),
                  TextSpan(
                      text: '                                   Download Mobile App - Android',
                      style: const TextStyle(color: Colors.blue),
                      recognizer: TapGestureRecognizer()
                        ..onTap = () {
                          launch('https://dublinbus-diy.site/download/DublinBus.apk');
                        }),
                  TextSpan(
                      text: '                                   Download Desktop App - Linux',
                      style: const TextStyle(color: Colors.blue),
                      recognizer: TapGestureRecognizer()
                        ..onTap = () {
                          launch('https://dublinbus-diy.site/download/DublinBus.deb');
                        }),
                ],
              ),
            ),
          ),
          Container(
            color: Colors.indigo[50],
            alignment: Alignment.center,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                Container(
                  child: RichText(
                    text: TextSpan(
                      children: [
                        const TextSpan(
                          text: ''' 
        \n\n\nCheck out Transport for Ireland below for more in-depth guide to public transport in Dublin. 
  There you wll find information about the capital???s trains, trams and buses all in one place, along with 
    a Dublin rail map that makes understanding how to get around in Dublin easier.
                                               ''',
                          style: TextStyle(color: Colors.black, fontSize: 15),
                        ),
                        TextSpan(
                            text: '           Visit the TFI website here\n',
                            style: const TextStyle(color: Colors.blue),
                            recognizer: TapGestureRecognizer()
                              ..onTap = () {
                                launch(
                                    'https://www.nationaltransport.ie/tfi-smarter-travel/public-transport/');
                              }),
                        const TextSpan(
                          text: ''' 
        \n\nYou can also check out the Leap Card link here. A TFI Leap Card is a prepaid travel card 
  that is the easiest way to pay your fare on public transport around Ireland.
        To get more information on Leap Card fares or to apply for a Leap card.
                                               ''',
                          style: const TextStyle(color: Colors.black, fontSize: 15),
                        ),
                        TextSpan(
                            text: '           Visit the Leap Card website here\n\n',
                            style: const TextStyle(color: Colors.blue),
                            recognizer: TapGestureRecognizer()
                              ..onTap = () {
                                launch('https://leapcard.ie/Home/index.html');
                              }),
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
